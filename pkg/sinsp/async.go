package sinsp

/*
#include <stdlib.h>
#include <stdint.h>
#include <stdbool.h>

typedef bool (*cb_wait_t)(void* wait_ctx);

typedef struct async_extractor_info
{
	uint64_t evtnum;
	uint32_t id;
	uint32_t ftype;
	char* arg;
	char* data;
	uint32_t datalen;
	uint32_t field_present;
	char* res;
	int32_t rc;
	cb_wait_t cb_wait;
	void* wait_ctx;
} async_extractor_info;

bool wait_bridge(async_extractor_info *info)
{
   return info->cb_wait(info->wait_ctx);
};
*/
import "C"
import "unsafe"

// RegisterAsyncExtractors is an helper function to be used within plugin_register_async_extractor.
//
// Intended usage as in the following example:
//
//     //export plugin_extract_str
//     func plugin_extract_str(pluginState unsafe.Pointer, evtnum uint64, id uint32, arg *byte, data *byte, datalen uint32) unsafe.Pointer {
//     	...
//     }
//
//     //export plugin_register_async_extractor
//     func plugin_register_async_extractor(pluginState unsafe.Pointer, asyncExtractorInfo unsafe.Pointer) int32 {
//     	return sinsp.RegisterAsyncExtractors(pluginState, asyncExtractorInfo, plugin_extract_str)
//     }
//
func RegisterAsyncExtractors(
	pluginState unsafe.Pointer,
	asyncExtractorInfo unsafe.Pointer,
	strExtractorFunc PluginExtractStrFunc,
) int32 {
	go func() {
		info := (*C.async_extractor_info)(asyncExtractorInfo)
		(*info).rc = C.int32_t(ScapSuccess)
		for C.wait_bridge(info) {
			switch uint32(info.ftype) {
			case ParamTypeCharBuf:
				if strExtractorFunc != nil {
					(*info).res = (*C.char)(unsafe.Pointer(strExtractorFunc(
						pluginState,
						uint64(info.evtnum),
						uint32(info.id),
						(*byte)(unsafe.Pointer(info.arg)),
						(*byte)(unsafe.Pointer(info.data)),
						uint32(info.datalen),
					)))
				} else {
					(*info).rc = C.int32_t(ScapNotSupported)
				}

			default:
				(*info).rc = C.int32_t(ScapNotSupported)
			}
		}
	}()
	return ScapSuccess
}
