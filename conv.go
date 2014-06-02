package rocksgo

// #include "rocksdb/c.h"
import "C"
import (
	"reflect"
	"unsafe"
)

func boolToUchar(b bool) C.uchar {
	uc := C.uchar(0)
	if b {
		uc = C.uchar(1)
	}
	return uc
}

func ucharToBool(uc C.uchar) bool {
	if uc == C.uchar(0) {
		return false
	}
	return true
}

// btoi converts a bool value to int
func btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}

// stringToChar returns *C.char from string
func stringToChar(s string) *C.char {
	ptrStr := (*reflect.StringHeader)(unsafe.Pointer(&s))

	return (*C.char)(unsafe.Pointer(ptrStr.Data))
}
