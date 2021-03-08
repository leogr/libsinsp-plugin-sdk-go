package main

import "C"
import (
	"encoding/json"
	"log"
	"unsafe"

	"github.com/ldegio/libsinsp-plugin-sdk-go/pkg/sinsp"
)

// Plugin consts
const (
	PluginID          uint32 = 1111
	PluginName               = "async"
	PluginDescription        = "async extractor example"
)

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
	return sinsp.TypeExtractorPlugin
}

//export plugin_init
func plugin_init(config *C.char, rc *int32) unsafe.Pointer {
	log.Printf("[%s] plugin_init\n", PluginName)
	log.Printf("config string:\n%s\n", C.GoString(config))

	*rc = sinsp.ScapSuccess

	return nil
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

//export plugin_get_fields
func plugin_get_fields() *C.char {
	log.Printf("[%s] plugin_get_fields\n", PluginName)
	flds := []sinsp.FieldEntry{
		{Type: "string", Name: "async.field", Desc: "TBD"},
	}

	b, err := json.Marshal(&flds)
	if err != nil {
		gLastError = err
		return nil
	}

	return C.CString(string(b))
}

//export plugin_extract_str
func plugin_extract_str(pluginState unsafe.Pointer, evtnum uint64, id uint32, arg *byte, data *byte, datalen uint32) *byte {
	//log.Printf("[%s] plugin_extract_str\n", PluginName)
	return (*byte)(unsafe.Pointer(C.CString("ciao")))
}

//export plugin_register_async_extractor
func plugin_register_async_extractor(pluginState unsafe.Pointer, asyncExtractorInfo unsafe.Pointer) int32 {
	log.Printf("[%s] plugin_register_async_extractor\n", PluginName)
	return sinsp.RegisterAsyncExtractors(pluginState, asyncExtractorInfo, plugin_extract_str)
}

func main() {}
