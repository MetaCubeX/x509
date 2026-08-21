// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build darwin

package macos

import (
	"errors"
	"strconv"
	"unsafe"
)

// Based on https://opensource.apple.com/source/Security/Security-59306.41.2/base/Security.h

const (
	// various macOS error codes that can be returned from
	// SecTrustEvaluateWithError that we can map to Go cert
	// verification error types.
	ErrSecCertificateExpired = -67818
	ErrSecHostNameMismatch   = -67602
	ErrSecNotTrusted         = -67843
)

type OSStatus struct {
	call   string
	status int32
}

func (s OSStatus) Error() string {
	return s.call + " error: " + strconv.Itoa(int(s.status))
}

//go:cgo_import_dynamic x509_SecTrustCreateWithCertificates SecTrustCreateWithCertificates "/System/Library/Frameworks/Security.framework/Versions/A/Security"

func SecTrustCreateWithCertificates(certs CFRef, policies CFRef) (CFRef, error) {
	var trustObj CFRef
	ret, _, _ := syscall_syscall6(x509_SecTrustCreateWithCertificates_trampoline_addr, uintptr(certs), uintptr(policies),
		uintptr(unsafe.Pointer(&trustObj)), 0, 0, 0)
	if int32(ret) != 0 {
		return 0, OSStatus{"SecTrustCreateWithCertificates", int32(ret)}
	}
	return trustObj, nil
}

var x509_SecTrustCreateWithCertificates_trampoline_addr uintptr

//go:cgo_import_dynamic x509_SecCertificateCreateWithData SecCertificateCreateWithData "/System/Library/Frameworks/Security.framework/Versions/A/Security"

func SecCertificateCreateWithData(b []byte) (CFRef, error) {
	data := BytesToCFData(b)
	defer CFRelease(data)
	ret, _, _ := syscall_syscall6(x509_SecCertificateCreateWithData_trampoline_addr, kCFAllocatorDefault, uintptr(data), 0, 0, 0, 0)
	// Returns NULL if the data passed in the data parameter is not a valid
	// DER-encoded X.509 certificate.
	if ret == 0 {
		return 0, errors.New("SecCertificateCreateWithData: invalid certificate")
	}
	return CFRef(ret), nil
}

var x509_SecCertificateCreateWithData_trampoline_addr uintptr

//go:cgo_import_dynamic x509_SecPolicyCreateSSL SecPolicyCreateSSL "/System/Library/Frameworks/Security.framework/Versions/A/Security"

func SecPolicyCreateSSL(name string) (CFRef, error) {
	var hostname CFString
	if name != "" {
		hostname = StringToCFString(name)
		defer CFRelease(CFRef(hostname))
	}
	ret, _, _ := syscall_syscall6(x509_SecPolicyCreateSSL_trampoline_addr, 1 /* true */, uintptr(hostname), 0, 0, 0, 0)
	if ret == 0 {
		return 0, OSStatus{"SecPolicyCreateSSL", int32(ret)}
	}
	return CFRef(ret), nil
}

var x509_SecPolicyCreateSSL_trampoline_addr uintptr

//go:cgo_import_dynamic x509_SecTrustSetVerifyDate SecTrustSetVerifyDate "/System/Library/Frameworks/Security.framework/Versions/A/Security"

func SecTrustSetVerifyDate(trustObj CFRef, dateRef CFRef) error {
	ret, _, _ := syscall_syscall6(x509_SecTrustSetVerifyDate_trampoline_addr, uintptr(trustObj), uintptr(dateRef), 0, 0, 0, 0)
	if int32(ret) != 0 {
		return OSStatus{"SecTrustSetVerifyDate", int32(ret)}
	}
	return nil
}

var x509_SecTrustSetVerifyDate_trampoline_addr uintptr

//go:cgo_import_dynamic x509_SecTrustEvaluate SecTrustEvaluate "/System/Library/Frameworks/Security.framework/Versions/A/Security"

func SecTrustEvaluate(trustObj CFRef) (CFRef, error) {
	var result CFRef
	ret, _, _ := syscall_syscall6(x509_SecTrustEvaluate_trampoline_addr, uintptr(trustObj), uintptr(unsafe.Pointer(&result)), 0, 0, 0, 0)
	if int32(ret) != 0 {
		return 0, OSStatus{"SecTrustEvaluate", int32(ret)}
	}
	return CFRef(result), nil
}

var x509_SecTrustEvaluate_trampoline_addr uintptr

//go:cgo_import_dynamic x509_SecTrustEvaluateWithError SecTrustEvaluateWithError "/System/Library/Frameworks/Security.framework/Versions/A/Security"

func SecTrustEvaluateWithError(trustObj CFRef) (int, error) {
	var errRef CFRef
	ret, _, _ := syscall_syscall6(x509_SecTrustEvaluateWithError_trampoline_addr, uintptr(trustObj), uintptr(unsafe.Pointer(&errRef)), 0, 0, 0, 0)
	if int32(ret) != 1 {
		errStr := CFErrorCopyDescription(errRef)
		err := errors.New(CFStringToString(errStr))
		errCode := CFErrorGetCode(errRef)
		CFRelease(errRef)
		CFRelease(errStr)
		return errCode, err
	}
	return 0, nil
}

var x509_SecTrustEvaluateWithError_trampoline_addr uintptr

//go:cgo_import_dynamic x509_SecTrustGetCertificateCount SecTrustGetCertificateCount "/System/Library/Frameworks/Security.framework/Versions/A/Security"

func SecTrustGetCertificateCount(trustObj CFRef) int {
	ret, _, _ := syscall_syscall6(x509_SecTrustGetCertificateCount_trampoline_addr, uintptr(trustObj), 0, 0, 0, 0, 0)
	return int(ret)
}

var x509_SecTrustGetCertificateCount_trampoline_addr uintptr

//go:cgo_import_dynamic x509_SecTrustGetCertificateAtIndex SecTrustGetCertificateAtIndex "/System/Library/Frameworks/Security.framework/Versions/A/Security"

func SecTrustGetCertificateAtIndex(trustObj CFRef, i int) (CFRef, error) {
	ret, _, _ := syscall_syscall6(x509_SecTrustGetCertificateAtIndex_trampoline_addr, uintptr(trustObj), uintptr(i), 0, 0, 0, 0)
	if ret == 0 {
		return 0, OSStatus{"SecTrustGetCertificateAtIndex", int32(ret)}
	}
	return CFRef(ret), nil
}

var x509_SecTrustGetCertificateAtIndex_trampoline_addr uintptr

//go:cgo_import_dynamic x509_SecCertificateCopyData SecCertificateCopyData "/System/Library/Frameworks/Security.framework/Versions/A/Security"

func SecCertificateCopyData(cert CFRef) ([]byte, error) {
	ret, _, _ := syscall_syscall6(x509_SecCertificateCopyData_trampoline_addr, uintptr(cert), 0, 0, 0, 0, 0)
	if ret == 0 {
		return nil, errors.New("x509: invalid certificate object")
	}
	b := CFDataToSlice(CFRef(ret))
	CFRelease(CFRef(ret))
	return b, nil
}

var x509_SecCertificateCopyData_trampoline_addr uintptr
