SHELL=/bin/bash -o pipefail

GO ?= go


.PHONY: examples/dummy
examples/dummy:
	GODEBUG=cgocheck=2 $(GO) build -buildmode=c-shared -o $@/libdummy.so $@/*.go

.PHONY: examples/async
examples/async:
	GODEBUG=cgocheck=2 $(GO) build -buildmode=c-shared -o $@/libasync.so $@/*.go

.PHONY: examples/batch
examples/batch:
	GODEBUG=cgocheck=2 $(GO) build -buildmode=c-shared -o $@/libbatch.so $@/*.go