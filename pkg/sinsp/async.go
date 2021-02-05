package sinsp

/*
#include <stdlib.h>
#include <stdint.h>
#include <stdbool.h>

typedef bool (*pfnWait)(void *waitCtx);

typedef struct async_extractor_info
{
	uint64_t evtnum;
	uint32_t id;
	char* arg;
	char* data;
	uint32_t datalen;
	uint32_t field_present;
	char* res;
	pfnWait wait;
	void *waitCtx;
} async_extractor_info;

bool wait_bridge(async_extractor_info *info)
{
   return info->wait(info->waitCtx);
};
*/
import "C"
import "unsafe"

func Wait(info unsafe.Pointer) bool {
	return bool(C.wait_bridge((*C.async_extractor_info)(info)))
}
