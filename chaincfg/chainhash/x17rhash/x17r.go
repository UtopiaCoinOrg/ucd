package x17r

/*
#cgo CFLAGS: -I./x17r_c_lib
#cgo LDFLAGS: -L./x17r_c_lib -lx17r
#include <stdio.h>
#include <stdlib.h>
#include "./x17r_c_lib/x17r.h"
*/
import "C"
import (
	"unsafe"
)

const X17HashSize = 32

func X17r_Sum256(input string) [X17HashSize]byte {
	var output [X17HashSize]byte
	in := (*C.uint8_t)(unsafe.Pointer(C.CString(input))) //(*C.char)(in)
	out := (*C.uint8_t)(C.malloc(X17HashSize))           //(*C.char)(out)
	C.x17r_hash(unsafe.Pointer(out), unsafe.Pointer(in), C.int(len(input)))
	p := uintptr(unsafe.Pointer(out))
	for i := 0; i < X17HashSize; i++ {
		value := *(*byte)(unsafe.Pointer(p))
		output[i] = value
		p += unsafe.Sizeof(value)
	}
	C.free(unsafe.Pointer(in))
	C.free(unsafe.Pointer(out))
	return output
}

func X17r_Sum256B(input string) []byte {
	var output [X17HashSize]byte
	in := (*C.uint8_t)(unsafe.Pointer(C.CString(input))) //(*C.char)(in)
	out := (*C.uint8_t)(C.malloc(X17HashSize))           //(*C.char)(out)
	C.x17r_hash(unsafe.Pointer(out), unsafe.Pointer(in), C.int(len(input)))
	p := uintptr(unsafe.Pointer(out))
	for i := 0; i < X17HashSize; i++ {
		value := *(*byte)(unsafe.Pointer(p))
		output[i] = value
		p += unsafe.Sizeof(value)
	}
	C.free(unsafe.Pointer(in))
	C.free(unsafe.Pointer(out))
	return output[:]
}
