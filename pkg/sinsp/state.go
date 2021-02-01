package sinsp

/*
#include <stdlib.h>
#include <stdint.h>

typedef struct {
   uint8_t* buf;
   uint32_t bufLen;
   void* goMem;
} state;
*/
import "C"
import "unsafe"

// NewStateContainer returns an opaque pointer to a memory container that
// may be safely passed back and forth to sinsp.
//
// A state container can allocate a byte buffer (sharable to C memory)
// and reference a Go pointer (suitable for a Go context).
// Both are persisted in memory until manually freed.
// A state container must be manually freed by using Free().
// It can be either used as the state of a source plugin or an open state of the source plugin.
func NewStateContainer() unsafe.Pointer {
	pCtx := (*C.state)(C.malloc(C.sizeof_state))
	pCtx.bufLen = 0
	pCtx.goMem = nil
	return unsafe.Pointer(pCtx)
}

// MakeBuffer allocates a C buffer of size s into p,
// assuming p is a state container created with NewStateContainer().
//
// If p contains a previously allocated buffer, it will be freed before creating the new one.
// Also, passing a zero size allows freeing a previously allocated buffer, if any.
func MakeBuffer(p unsafe.Pointer, s uint32) {
	state := (*C.state)(p)

	if state.bufLen > 0 {
		C.free(unsafe.Pointer(state.buf))
	}

	if s > 0 {
		state.buf = (*C.uint8_t)(C.malloc(C.ulong(s)))
	}

	state.bufLen = C.uint32_t(s)
}

// CopyToBuffer copies bytes from a b into the C buffer belonging to p,
// assuming p is a state container created with NewStateContainer().
//
// The buffer must be previously created with MakeBuffer().
// It returns the number of bytes copied, which will be the minimum of the buffer's size and len(b).
func CopyToBuffer(p unsafe.Pointer, b []byte) uint32 {
	state := (*C.state)(p)
	return uint32(copy((*[1 << 30]byte)(unsafe.Pointer(state.buf))[:int(state.bufLen):int(state.bufLen)][:], b))
}

// Buffer returns a pointer to the first element of the C buffer belonging to p, if any,
// assuming p is a state container created with NewStateContainer().
func Buffer(p unsafe.Pointer) *byte {
	return (*byte)((*C.state)(p).buf)
}

// SetContext sets the given reference ctx (a pointer to a Go allocated memory) into p,
// assuming p is a state container created with NewStateContainer().
//
// A previously set reference, if any, is removed from p, making it suitable for garbage collecting.
func SetContext(p unsafe.Pointer, ctx unsafe.Pointer) {
	state := (*C.state)(p)

	if state.goMem != nil {
		peristentPtrs.Delete(state.goMem)
	}

	state.goMem = ctx

	if ctx != nil {
		peristentPtrs.Store(ctx, ctx)
	}
}

// Context returns a pointer to Go allocated memory, if any, previously assigned into p with SetContext(),
// assuming p is a state container created with NewStateContainer().
func Context(p unsafe.Pointer) unsafe.Pointer {
	return (*C.state)(p).goMem
}

// Free disposes of any C and Go memory assigned to p and finally free P,
// assuming p is a state container created with NewStateContainer().
func Free(p unsafe.Pointer) {
	MakeBuffer(p, 0)
	SetContext(p, nil)
	C.free(p)
}
