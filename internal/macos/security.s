// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build darwin

#include "textflag.h"

TEXT x509_SecTrustCreateWithCertificates_trampoline<>(SB),NOSPLIT,$0-0
	JMP	x509_SecTrustCreateWithCertificates(SB)
GLOBL	·x509_SecTrustCreateWithCertificates_trampoline_addr(SB), RODATA, $8
DATA	·x509_SecTrustCreateWithCertificates_trampoline_addr(SB)/8, $x509_SecTrustCreateWithCertificates_trampoline<>(SB)

TEXT x509_SecCertificateCreateWithData_trampoline<>(SB),NOSPLIT,$0-0
	JMP	x509_SecCertificateCreateWithData(SB)
GLOBL	·x509_SecCertificateCreateWithData_trampoline_addr(SB), RODATA, $8
DATA	·x509_SecCertificateCreateWithData_trampoline_addr(SB)/8, $x509_SecCertificateCreateWithData_trampoline<>(SB)

TEXT x509_SecPolicyCreateSSL_trampoline<>(SB),NOSPLIT,$0-0
	JMP	x509_SecPolicyCreateSSL(SB)
GLOBL	·x509_SecPolicyCreateSSL_trampoline_addr(SB), RODATA, $8
DATA	·x509_SecPolicyCreateSSL_trampoline_addr(SB)/8, $x509_SecPolicyCreateSSL_trampoline<>(SB)

TEXT x509_SecTrustSetVerifyDate_trampoline<>(SB),NOSPLIT,$0-0
	JMP	x509_SecTrustSetVerifyDate(SB)
GLOBL	·x509_SecTrustSetVerifyDate_trampoline_addr(SB), RODATA, $8
DATA	·x509_SecTrustSetVerifyDate_trampoline_addr(SB)/8, $x509_SecTrustSetVerifyDate_trampoline<>(SB)

TEXT x509_SecTrustEvaluate_trampoline<>(SB),NOSPLIT,$0-0
	JMP	x509_SecTrustEvaluate(SB)
GLOBL	·x509_SecTrustEvaluate_trampoline_addr(SB), RODATA, $8
DATA	·x509_SecTrustEvaluate_trampoline_addr(SB)/8, $x509_SecTrustEvaluate_trampoline<>(SB)

TEXT x509_SecTrustEvaluateWithError_trampoline<>(SB),NOSPLIT,$0-0
	JMP	x509_SecTrustEvaluateWithError(SB)
GLOBL	·x509_SecTrustEvaluateWithError_trampoline_addr(SB), RODATA, $8
DATA	·x509_SecTrustEvaluateWithError_trampoline_addr(SB)/8, $x509_SecTrustEvaluateWithError_trampoline<>(SB)

TEXT x509_SecTrustGetCertificateCount_trampoline<>(SB),NOSPLIT,$0-0
	JMP	x509_SecTrustGetCertificateCount(SB)
GLOBL	·x509_SecTrustGetCertificateCount_trampoline_addr(SB), RODATA, $8
DATA	·x509_SecTrustGetCertificateCount_trampoline_addr(SB)/8, $x509_SecTrustGetCertificateCount_trampoline<>(SB)

TEXT x509_SecTrustGetCertificateAtIndex_trampoline<>(SB),NOSPLIT,$0-0
	JMP	x509_SecTrustGetCertificateAtIndex(SB)
GLOBL	·x509_SecTrustGetCertificateAtIndex_trampoline_addr(SB), RODATA, $8
DATA	·x509_SecTrustGetCertificateAtIndex_trampoline_addr(SB)/8, $x509_SecTrustGetCertificateAtIndex_trampoline<>(SB)

TEXT x509_SecCertificateCopyData_trampoline<>(SB),NOSPLIT,$0-0
	JMP	x509_SecCertificateCopyData(SB)
GLOBL	·x509_SecCertificateCopyData_trampoline_addr(SB), RODATA, $8
DATA	·x509_SecCertificateCopyData_trampoline_addr(SB)/8, $x509_SecCertificateCopyData_trampoline<>(SB)
