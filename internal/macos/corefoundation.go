// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build darwin

// Package macos provides cgo-less wrappers for Core Foundation and
// Security.framework, similarly to how package syscall provides access to
// libSystem.dylib.
package macos

import (
	"bytes"
	"errors"
	"math"
	"runtime"
	"time"
	"unsafe"
)

// CFRef is an opaque reference to a Core Foundation object. It is a pointer,
// but to memory not owned by Go, so not an unsafe.Pointer.
type CFRef uintptr

// CFDataToSlice returns a copy of the contents of data as a bytes slice.
func CFDataToSlice(data CFRef) []byte {
	length := CFDataGetLength(data)
	ptr := CFDataGetBytePtr(data)
	src := unsafe.Slice((*byte)(unsafe.Pointer(ptr)), length)
	return bytes.Clone(src)
}

// CFStringToString returns a Go string representation of the passed
// in CFString, or an empty string if it's invalid.
func CFStringToString(ref CFRef) string {
	data, err := CFStringCreateExternalRepresentation(ref)
	if err != nil {
		return ""
	}
	b := CFDataToSlice(data)
	CFRelease(data)
	return string(b)
}

// TimeToCFDateRef converts a time.Time into an apple CFDateRef.
func TimeToCFDateRef(t time.Time) CFRef {
	secs := t.Sub(time.Date(2001, 1, 1, 0, 0, 0, 0, time.UTC)).Seconds()
	ref := CFDateCreate(secs)
	return ref
}

type CFString CFRef

const kCFAllocatorDefault = 0
const kCFStringEncodingUTF8 = 0x08000100

//go:cgo_import_dynamic x509_CFDataCreate CFDataCreate "/System/Library/Frameworks/CoreFoundation.framework/Versions/A/CoreFoundation"

func BytesToCFData(b []byte) CFRef {
	p := unsafe.Pointer(unsafe.SliceData(b))
	ret, _, _ := syscall_syscall6(x509_CFDataCreate_trampoline_addr, kCFAllocatorDefault, uintptr(p), uintptr(len(b)), 0, 0, 0)
	runtime.KeepAlive(p)
	return CFRef(ret)
}

var x509_CFDataCreate_trampoline_addr uintptr

//go:cgo_import_dynamic x509_CFStringCreateWithBytes CFStringCreateWithBytes "/System/Library/Frameworks/CoreFoundation.framework/Versions/A/CoreFoundation"

// StringToCFString returns a copy of the UTF-8 contents of s as a new CFString.
func StringToCFString(s string) CFString {
	p := unsafe.Pointer(unsafe.StringData(s))
	ret, _, _ := syscall_syscall6(x509_CFStringCreateWithBytes_trampoline_addr, kCFAllocatorDefault, uintptr(p),
		uintptr(len(s)), uintptr(kCFStringEncodingUTF8), 0 /* isExternalRepresentation */, 0)
	runtime.KeepAlive(p)
	return CFString(ret)
}

var x509_CFStringCreateWithBytes_trampoline_addr uintptr

//go:cgo_import_dynamic x509_CFDictionaryGetValueIfPresent CFDictionaryGetValueIfPresent "/System/Library/Frameworks/CoreFoundation.framework/Versions/A/CoreFoundation"

func CFDictionaryGetValueIfPresent(dict CFRef, key CFString) (value CFRef, ok bool) {
	ret, _, _ := syscall_syscall6(x509_CFDictionaryGetValueIfPresent_trampoline_addr, uintptr(dict), uintptr(key),
		uintptr(unsafe.Pointer(&value)), 0, 0, 0)
	if ret == 0 {
		return 0, false
	}
	return value, true
}

var x509_CFDictionaryGetValueIfPresent_trampoline_addr uintptr

const kCFNumberSInt32Type = 3

//go:cgo_import_dynamic x509_CFNumberGetValue CFNumberGetValue "/System/Library/Frameworks/CoreFoundation.framework/Versions/A/CoreFoundation"

func CFNumberGetValue(num CFRef) (int32, error) {
	var value int32
	ret, _, _ := syscall_syscall6(x509_CFNumberGetValue_trampoline_addr, uintptr(num), uintptr(kCFNumberSInt32Type),
		uintptr(unsafe.Pointer(&value)), 0, 0, 0)
	if ret == 0 {
		return 0, errors.New("CFNumberGetValue call failed")
	}
	return value, nil
}

var x509_CFNumberGetValue_trampoline_addr uintptr

//go:cgo_import_dynamic x509_CFDataGetLength CFDataGetLength "/System/Library/Frameworks/CoreFoundation.framework/Versions/A/CoreFoundation"

func CFDataGetLength(data CFRef) int {
	ret, _, _ := syscall_syscall6(x509_CFDataGetLength_trampoline_addr, uintptr(data), 0, 0, 0, 0, 0)
	return int(ret)
}

var x509_CFDataGetLength_trampoline_addr uintptr

//go:cgo_import_dynamic x509_CFDataGetBytePtr CFDataGetBytePtr "/System/Library/Frameworks/CoreFoundation.framework/Versions/A/CoreFoundation"

func CFDataGetBytePtr(data CFRef) uintptr {
	ret, _, _ := syscall_syscall6(x509_CFDataGetBytePtr_trampoline_addr, uintptr(data), 0, 0, 0, 0, 0)
	return ret
}

var x509_CFDataGetBytePtr_trampoline_addr uintptr

//go:cgo_import_dynamic x509_CFArrayGetCount CFArrayGetCount "/System/Library/Frameworks/CoreFoundation.framework/Versions/A/CoreFoundation"

func CFArrayGetCount(array CFRef) int {
	ret, _, _ := syscall_syscall6(x509_CFArrayGetCount_trampoline_addr, uintptr(array), 0, 0, 0, 0, 0)
	return int(ret)
}

var x509_CFArrayGetCount_trampoline_addr uintptr

//go:cgo_import_dynamic x509_CFArrayGetValueAtIndex CFArrayGetValueAtIndex "/System/Library/Frameworks/CoreFoundation.framework/Versions/A/CoreFoundation"

func CFArrayGetValueAtIndex(array CFRef, index int) CFRef {
	ret, _, _ := syscall_syscall6(x509_CFArrayGetValueAtIndex_trampoline_addr, uintptr(array), uintptr(index), 0, 0, 0, 0)
	return CFRef(ret)
}

var x509_CFArrayGetValueAtIndex_trampoline_addr uintptr

//go:cgo_import_dynamic x509_CFEqual CFEqual "/System/Library/Frameworks/CoreFoundation.framework/Versions/A/CoreFoundation"

func CFEqual(a, b CFRef) bool {
	ret, _, _ := syscall_syscall6(x509_CFEqual_trampoline_addr, uintptr(a), uintptr(b), 0, 0, 0, 0)
	return ret == 1
}

var x509_CFEqual_trampoline_addr uintptr

//go:cgo_import_dynamic x509_CFRelease CFRelease "/System/Library/Frameworks/CoreFoundation.framework/Versions/A/CoreFoundation"

func CFRelease(ref CFRef) {
	syscall_syscall6(x509_CFRelease_trampoline_addr, uintptr(ref), 0, 0, 0, 0, 0)
}

var x509_CFRelease_trampoline_addr uintptr

//go:cgo_import_dynamic x509_CFArrayCreateMutable CFArrayCreateMutable "/System/Library/Frameworks/CoreFoundation.framework/Versions/A/CoreFoundation"

func CFArrayCreateMutable() CFRef {
	ret, _, _ := syscall_syscall6(x509_CFArrayCreateMutable_trampoline_addr, kCFAllocatorDefault, 0, 0 /* kCFTypeArrayCallBacks */, 0, 0, 0)
	return CFRef(ret)
}

var x509_CFArrayCreateMutable_trampoline_addr uintptr

//go:cgo_import_dynamic x509_CFArrayAppendValue CFArrayAppendValue "/System/Library/Frameworks/CoreFoundation.framework/Versions/A/CoreFoundation"

func CFArrayAppendValue(array CFRef, val CFRef) {
	syscall_syscall6(x509_CFArrayAppendValue_trampoline_addr, uintptr(array), uintptr(val), 0, 0, 0, 0)
}

var x509_CFArrayAppendValue_trampoline_addr uintptr

//go:cgo_import_dynamic x509_CFDateCreate CFDateCreate "/System/Library/Frameworks/CoreFoundation.framework/Versions/A/CoreFoundation"

func CFDateCreate(seconds float64) CFRef {
	ret, _, _ := syscall_syscall6(x509_CFDateCreate_trampoline_addr, kCFAllocatorDefault, uintptr(math.Float64bits(seconds)), 0, 0, 0, 0)
	return CFRef(ret)
}

var x509_CFDateCreate_trampoline_addr uintptr

//go:cgo_import_dynamic x509_CFErrorCopyDescription CFErrorCopyDescription "/System/Library/Frameworks/CoreFoundation.framework/Versions/A/CoreFoundation"

func CFErrorCopyDescription(errRef CFRef) CFRef {
	ret, _, _ := syscall_syscall6(x509_CFErrorCopyDescription_trampoline_addr, uintptr(errRef), 0, 0, 0, 0, 0)
	return CFRef(ret)
}

var x509_CFErrorCopyDescription_trampoline_addr uintptr

//go:cgo_import_dynamic x509_CFErrorGetCode CFErrorGetCode "/System/Library/Frameworks/CoreFoundation.framework/Versions/A/CoreFoundation"

func CFErrorGetCode(errRef CFRef) int {
	ret, _, _ := syscall_syscall6(x509_CFErrorGetCode_trampoline_addr, uintptr(errRef), 0, 0, 0, 0, 0)
	return int(ret)
}

var x509_CFErrorGetCode_trampoline_addr uintptr

//go:cgo_import_dynamic x509_CFStringCreateExternalRepresentation CFStringCreateExternalRepresentation "/System/Library/Frameworks/CoreFoundation.framework/Versions/A/CoreFoundation"

func CFStringCreateExternalRepresentation(strRef CFRef) (CFRef, error) {
	ret, _, _ := syscall_syscall6(x509_CFStringCreateExternalRepresentation_trampoline_addr, kCFAllocatorDefault, uintptr(strRef), kCFStringEncodingUTF8, 0, 0, 0)
	if ret == 0 {
		return 0, errors.New("string can't be represented as UTF-8")
	}
	return CFRef(ret), nil
}

var x509_CFStringCreateExternalRepresentation_trampoline_addr uintptr

// ReleaseCFArray iterates through an array, releasing its contents, and then
// releases the array itself. This is necessary because we cannot, easily, set the
// CFArrayCallBacks argument when creating CFArrays.
func ReleaseCFArray(array CFRef) {
	for i := 0; i < CFArrayGetCount(array); i++ {
		ref := CFArrayGetValueAtIndex(array, i)
		CFRelease(ref)
	}
	CFRelease(array)
}
