package sinsp

import "C"
import "unsafe"

// PluginExtractStrFunc represents one common signature for the implementation of the `plugin_event_to_string()`
type PluginExtractStrFunc func(pluginState unsafe.Pointer, evtnum uint64, id uint32, arg *byte, data *byte, datalen uint32) *byte
