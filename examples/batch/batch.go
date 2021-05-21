package main

/*
#include <stdlib.h>
#include <stdint.h>
*/
import "C"
import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"unsafe"

	"github.com/ldegio/libsinsp-plugin-sdk-go/pkg/sinsp"
)

// Plugin consts
const (
	PluginID          uint32 = 111
	PluginName               = "batch"
	PluginDescription        = "do almost nothing"
)

const nextBufSize uint32 = 65535
const outBufSize uint32 = 4096

///////////////////////////////////////////////////////////////////////////////

type pluginCtx struct {
	m       map[int]string
	counter int
}

// todo: plugin_get_last_error() needs context as argument to avoid having this global
var gLastError error

//export plugin_get_type
func plugin_get_type() uint32 {
	log.Printf("[%s] plugin_get_type\n", PluginName)
	return sinsp.TypeSourcePlugin
}

//export plugin_init
func plugin_init(config *C.char, rc *int32) unsafe.Pointer {
	log.Printf("[%s] plugin_init\n", PluginName)
	log.Printf("config string:\n%s\n", C.GoString(config))

	pState := sinsp.NewStateContainer()
	sinsp.MakeBuffer(pState, outBufSize)
	*rc = sinsp.ScapSuccess

	return pState
}

//export plugin_get_last_error
func plugin_get_last_error() *C.char {
	log.Printf("[%s] plugin_get_last_error\n", PluginName)
	if gLastError != nil {
		return C.CString(gLastError.Error())
	}
	return nil
}

//export plugin_destroy
func plugin_destroy(pState unsafe.Pointer) {
	log.Printf("[%s] plugin_destroy\n", PluginName)
	sinsp.Free(pState)
}

//export plugin_get_id
func plugin_get_id() uint32 {
	log.Printf("[%s] plugin_get_id\n", PluginName)
	return PluginID
}

//export plugin_get_name
func plugin_get_name() *C.char {
	log.Printf("[%s] plugin_get_name\n", PluginName)
	return C.CString(PluginName)
}

//export plugin_get_description
func plugin_get_description() *C.char {
	log.Printf("[%s] plugin_get_description\n", PluginName)
	return C.CString(PluginDescription)
}

// export plugin_get_required_api_version
func plugin_get_required_api_version() *C.char {
	return C.CString("1.0.0")
}

//export plugin_get_fields
func plugin_get_fields() *C.char {
	log.Printf("[%s] plugin_get_fields\n", PluginName)
	flds := []sinsp.FieldEntry{
		{Type: "string", Name: "dummy.count", Desc: "TBD"},
	}

	b, err := json.Marshal(&flds)
	if err != nil {
		gLastError = err
		return nil
	}

	return C.CString(string(b))
}

//export plugin_open
func plugin_open(pState unsafe.Pointer, params *C.char, rc *int32) unsafe.Pointer {
	input := C.GoString(params)
	log.Printf("[%s] plugin_open, params: %s\n", PluginName, input)

	m := &pluginCtx{}
	m.m = make(map[int]string)
	m.m[4] = "ciao"

	oState := sinsp.NewStateContainer()
	sinsp.MakeBuffer(oState, nextBufSize)
	sinsp.SetContext(oState, unsafe.Pointer(m))

	*rc = sinsp.ScapSuccess
	return oState
}

//export plugin_close
func plugin_close(pState unsafe.Pointer, oState unsafe.Pointer) {
	log.Printf("[%s] plugin_close\n", PluginName)
	m := (*pluginCtx)(sinsp.Context(oState))
	log.Printf("[%s] Dump context before freeing\n", PluginName)
	fmt.Println(m)
	sinsp.Free(oState)
}

func next(plgState unsafe.Pointer, oState unsafe.Pointer, data *[]byte, ts *uint64) int32 {

	m := (*pluginCtx)(sinsp.Context(oState))

	// dummy plugin always produce "dummy" data
	dummy := fmt.Sprintf("dummy%d", int(m.counter))
	m.counter++

	// Put something not usefull in Go memory
	m.m[rand.Intn(100)] = dummy

	bdummy := []byte(dummy)
	data = &bdummy

	return sinsp.ScapSuccess
}

//export plugin_next
func plugin_next(pState unsafe.Pointer, oState unsafe.Pointer, data **byte, datalen *uint32, ts *uint64) int32 {
	var nextData []byte

	res := next(pState, oState, &nextData, ts)
	if res == sinsp.ScapSuccess {
		// Copy to and return the event buffer
		*datalen = sinsp.CopyToBuffer(oState, nextData)
		*data = sinsp.Buffer(oState)
	}

	return res
}

//export plugin_event_to_string
func plugin_event_to_string(data *C.char, datalen uint32) *C.char {
	log.Printf("[%s] plugin_event_to_string\n", PluginName)
	// do something dummy with the string
	s := fmt.Sprintf("evt-to-string(len=%d): %s", datalen, C.GoStringN(data, C.int(datalen)))
	return C.CString(s)
}

//export plugin_next_batch
func plugin_next_batch(pState unsafe.Pointer, oState unsafe.Pointer, data **byte, datalen *uint32) int32 {
	return sinsp.NextBatch(pState, oState, data, datalen, next)
}

func main() {}
