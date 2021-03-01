package sinsp

// SCAP return types
const (
	ScapSuccess         int32 = 0
	ScapFailure         int32 = 1
	ScapTimeout         int32 = -1
	ScapIllegalInput    int32 = 3
	ScapNotFound        int32 = 4
	ScapInputTooSmall   int32 = 5
	ScapEOF             int32 = 6
	ScapUnexpectedBlock int32 = 7
	ScapVersionMismatch int32 = 8
	ScapNotSupported    int32 = 9
)

// Plugin types
const (
	TypeSourcePlugin    uint32 = 1
	TypeExtractorPlugin uint32 = 2
)

const MaxEvtSize uint32 = 65635
const MaxNextBufSize uint32 = 4 * 1024 * 1024
