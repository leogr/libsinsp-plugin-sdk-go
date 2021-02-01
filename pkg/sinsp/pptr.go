package sinsp

import (
	"sync"
)

var peristentPtrs = &sync.Map{}
