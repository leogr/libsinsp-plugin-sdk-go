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

const (
	ParamTypeNone             uint32 = 0
	ParamTypeInt8             uint32 = 1
	ParamTypeInt16            uint32 = 2
	ParamTypeInt32            uint32 = 3
	ParamTypeInt64            uint32 = 4
	ParamTypeUintT8           uint32 = 5
	ParamTypeUint16           uint32 = 6
	ParamTypeUint32           uint32 = 7
	ParamTypeUint64           uint32 = 8
	ParamTypeCharBuf          uint32 = 9  // A printable buffer of bytes, NULL terminated
	ParamTypeByteBuf          uint32 = 10 // A raw buffer of bytes not suitable for printing
	ParamTypeErrno            uint32 = 11 // this is an INT64, but will be interpreted as an error code
	ParamTypeSockaddr         uint32 = 12 // A sockaddr structure, 1byte family + data
	ParamTypeSocktuple        uint32 = 13 // A sockaddr tuple,1byte family + 12byte data + 12byte data
	ParamTypeFd               uint32 = 14 // An fd, 64bit
	ParamTypePid              uint32 = 15 // A pid/tid, 64bit
	ParamTypeFdlist           uint32 = 16 // A list of fds, 16bit count + count * (64bit fd + 16bit flags)
	ParamTypeFspath           uint32 = 17 // A string containing a relative or absolute file system path, null terminated
	ParamTypeSyscallId        uint32 = 18 // A 16bit system call ID. Can be used as a key for the g_syscall_info_table table.
	ParamTypeSigYype          uint32 = 19 // An 8bit signal number
	ParamTypeRelTime          uint32 = 20 // A relative time. Seconds * 10^9  + nanoseconds. 64bit.
	ParamTypeAbsTime          uint32 = 21 // An absolute time interval. Seconds from epoch * 10^9  + nanoseconds. 64bit.
	ParamTypePort             uint32 = 22 // A TCP/UDP prt. 2 bytes.
	ParamTypeL4Proto          uint32 = 23 // A 1 byte IP protocol type.
	ParamTypeSockfamily       uint32 = 24 // A 1 byte socket family.
	ParamTypeBool             uint32 = 25 // A boolean value, 4 bytes.
	ParamTypeIpv4Addr         uint32 = 26 // A 4 byte raw IPv4 address.
	ParamTypeDyn              uint32 = 27 // Type can vary depending on the context. Used for filter fields like evt.rawarg.
	ParamTypeFlags8           uint32 = 28 // this is an UINT8, but will be interpreted as 8 bit flags.
	ParamTypeFlags16          uint32 = 29 // this is an UINT16, but will be interpreted as 16 bit flags.
	ParamTypeFlags32          uint32 = 30 // this is an UINT32, but will be interpreted as 32 bit flags.
	ParamTypeUid              uint32 = 31 // this is an UINT32, MAX_UINT32 will be interpreted as no value.
	ParamTypeGid              uint32 = 32 // this is an UINT32, MAX_UINT32 will be interpreted as no value.
	ParamTypeDouble           uint32 = 33 // this is a double precision floating point number.
	ParamTypeSigSet           uint32 = 34 // sigset_t. I only store the lower UINT32 of it
	ParamTypeCharBufArray     uint32 = 35 // Pointer to an array of strings, exported by the user events decoder. 64bit. For internal use only.
	ParamTypeCharBufPairArray uint32 = 36 // Pointer to an array of string pairs, exported by the user events decoder. 64bit. For internal use only.
	ParamTypeIpv4Net          uint32 = 37 // An IPv4 network.
	ParamTypeIpv6Addr         uint32 = 38 // A 16 byte raw IPv6 address.
	ParamTypeIpv6Net          uint32 = 39 // An IPv6 network.
	ParamTypeIpAddr           uint32 = 40 // Either an IPv4 or IPv6 address. The length indicates which one it is.
	ParamTypeIpNet            uint32 = 41 // Either an IPv4 or IPv6 network. The length indicates which one it is.
	ParamTypeMode             uint32 = 42 // a 32 bit bitmask to represent file modes.
	ParamTypeFsRelPath        uint32 = 43 // A path relative to a dirfd.
	ParamTypeMax              uint32 = 44 // array size
)
