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
	char* arg;
	char* data;
	uint32_t datalen;
	uint32_t field_present;
	char* res;
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

func Wait(info unsafe.Pointer) bool {
	return bool(C.wait_bridge((*C.async_extractor_info)(info)))
}
