package mylib

/*
#cgo LDFLAGS: -L. -L./mylib/ -lmylib

#include "mylib.h"
*/
import "C"

//import "unsafe"

// Go wrapper function that calls a function from the C library
func MyFunction(input string) int {
	//cInput := C.CString(input)
	//defer C.free(unsafe.Pointer(cInput))

	//cOutput := C.MyCFunction(cInput)
	return int(C.SayHello())
	//defer C.free(unsafe.Pointer(cOutput))

	//return input
	//C.GoString(cOutput)
}
