package sinsp

// SCAP return types
const (
	ScapSuccess int32 = 0
	ScapFailure int32 = 1
	ScapTimeout int32 = -1
)

// Plugin types
const (
	TypeSourcePlugin    uint32 = 1
	TypeExtractorPlugin uint32 = 2
)
