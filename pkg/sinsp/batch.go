package sinsp

import (
	"encoding/binary"
	"unsafe"
)

type batchContext struct {
	nextBatchLastTs   uint64
	nextBatchLastData []byte
}

// NextFunc is the function type required by NextBatch().
type NextFunc func(plgState unsafe.Pointer, openState unsafe.Pointer, data *[]byte, ts *uint64) int32

// NextBatch is an helper function to be used within plugin_next_batch.
func NextBatch(plgState unsafe.Pointer, openState unsafe.Pointer, data **byte, datalen *uint32, nextf NextFunc) int32 {
	var ts uint64
	tsbuf := make([]byte, int(unsafe.Sizeof(ts)))
	var elen uint32
	elenbuf := make([]byte, int(unsafe.Sizeof(elen)))
	res := ScapSuccess
	*datalen = 0
	var pos uint32 = 0
	var nextData []byte

	bCtx := getBatchCtx(openState)
	if bCtx.nextBatchLastData != nil {
		//
		// There is leftover data from the previous call, copy it at the start
		// of the buffer
		//
		loData := &bCtx.nextBatchLastData
		binary.LittleEndian.PutUint64(tsbuf, bCtx.nextBatchLastTs)
		binary.LittleEndian.PutUint32(elenbuf, uint32(len(*loData)))
		pos += CopyToBufferAt(openState, tsbuf, pos)
		pos += CopyToBufferAt(openState, elenbuf, pos)
		pos += CopyToBufferAt(openState, *loData, pos)
	}

	bCtx.nextBatchLastData = nil

	for true {
		res = nextf(plgState, openState, &nextData, &ts)
		if res == ScapSuccess {
			endPos := pos + uint32(len(nextData)) + 12
			if endPos < MaxNextBufSize {
				// Copy the event into the buffer
				binary.LittleEndian.PutUint64(tsbuf, ts)
				binary.LittleEndian.PutUint32(elenbuf, uint32(len(nextData)))
				pos += CopyToBufferAt(openState, tsbuf, pos)
				pos += CopyToBufferAt(openState, elenbuf, pos)
				pos += CopyToBufferAt(openState, nextData, pos)
			} else {
				if pos > 0 {
					// Buffer full. Save this event for the next read
					bCtx.nextBatchLastTs = ts
					bCtx.nextBatchLastData = nextData
				} else {
					// This event is too big to fit in the buffer by itself.
					// Skip it.
					res = ScapTimeout
				}
				break
			}
		} else {
			break
		}
	}

	*data = Buffer(openState)
	*datalen = pos

	return res
}
