package x19r

/*
#cgo CFLAGS: -I./x19r_c_lib
#cgo LDFLAGS: -L./x19r_c_lib -lx19r
#include <stdio.h>
#include <stdlib.h>
#include "./x19r_c_lib/x19r.h"
*/
import "C"
import (
	"unsafe"
)

const X19HashSize = 32

func X19r_Sum256(input string) [X19HashSize]byte {
	var output [X19HashSize]byte
	in := (*C.uint8_t)(unsafe.Pointer(C.CString(input))) //(*C.char)(in)
	out := (*C.uint8_t)(C.malloc(X19HashSize))           //(*C.char)(out)
	C.x19r_hash(unsafe.Pointer(out), unsafe.Pointer(in), C.int(len(input)))
	p := uintptr(unsafe.Pointer(out))
	for i := 0; i < X19HashSize; i++ {
		value := *(*byte)(unsafe.Pointer(p))
		output[i] = value
		p += unsafe.Sizeof(value)
	}
	C.free(unsafe.Pointer(in))
	C.free(unsafe.Pointer(out))
	return output
}

func X19r_Sum256B(input string) []byte {
	var output [X19HashSize]byte
	in := (*C.uint8_t)(unsafe.Pointer(C.CString(input))) //(*C.char)(in)
	out := (*C.uint8_t)(C.malloc(X19HashSize))           //(*C.char)(out)
	C.x19r_hash(unsafe.Pointer(out), unsafe.Pointer(in), C.int(len(input)))
	p := uintptr(unsafe.Pointer(out))
	for i := 0; i < X19HashSize; i++ {
		value := *(*byte)(unsafe.Pointer(p))
		output[i] = value
		p += unsafe.Sizeof(value)
	}
	C.free(unsafe.Pointer(in))
	C.free(unsafe.Pointer(out))
	return output[:]
}
