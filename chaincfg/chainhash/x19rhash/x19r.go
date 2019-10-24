package x19r

/*
#cgo CFLAGS: -I./x19r_c_lib
#cgo LDFLAGS: -L./x19r_c_lib -I./scratch-x19r/arm64 -I./scratch-x19r/armv7 -I./scratch-x19r/armv7s -I./scratch-x19r/i386 -I./scratch-x19r/x86_64 -I./local/arm64-v8a -I./local/armeabi-v7a -I./local/x86 -I./local/x86_64 -lx19r
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
