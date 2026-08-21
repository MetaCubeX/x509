// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build darwin

#include "textflag.h"

// Keep the trampoline addresses in data symbols, as golang.org/x/sys/unix
// does, so Go code does not need internal/abi.FuncPCABI0.

TEXT x509_CFArrayGetCount_trampoline<>(SB),NOSPLIT,$0-0
	JMP	x509_CFArrayGetCount(SB)
GLOBL	·x509_CFArrayGetCount_trampoline_addr(SB), RODATA, $8
DATA	·x509_CFArrayGetCount_trampoline_addr(SB)/8, $x509_CFArrayGetCount_trampoline<>(SB)

TEXT x509_CFArrayGetValueAtIndex_trampoline<>(SB),NOSPLIT,$0-0
	JMP	x509_CFArrayGetValueAtIndex(SB)
GLOBL	·x509_CFArrayGetValueAtIndex_trampoline_addr(SB), RODATA, $8
DATA	·x509_CFArrayGetValueAtIndex_trampoline_addr(SB)/8, $x509_CFArrayGetValueAtIndex_trampoline<>(SB)

TEXT x509_CFDataGetBytePtr_trampoline<>(SB),NOSPLIT,$0-0
	JMP	x509_CFDataGetBytePtr(SB)
GLOBL	·x509_CFDataGetBytePtr_trampoline_addr(SB), RODATA, $8
DATA	·x509_CFDataGetBytePtr_trampoline_addr(SB)/8, $x509_CFDataGetBytePtr_trampoline<>(SB)

TEXT x509_CFDataGetLength_trampoline<>(SB),NOSPLIT,$0-0
	JMP	x509_CFDataGetLength(SB)
GLOBL	·x509_CFDataGetLength_trampoline_addr(SB), RODATA, $8
DATA	·x509_CFDataGetLength_trampoline_addr(SB)/8, $x509_CFDataGetLength_trampoline<>(SB)

TEXT x509_CFStringCreateWithBytes_trampoline<>(SB),NOSPLIT,$0-0
	JMP	x509_CFStringCreateWithBytes(SB)
GLOBL	·x509_CFStringCreateWithBytes_trampoline_addr(SB), RODATA, $8
DATA	·x509_CFStringCreateWithBytes_trampoline_addr(SB)/8, $x509_CFStringCreateWithBytes_trampoline<>(SB)

TEXT x509_CFRelease_trampoline<>(SB),NOSPLIT,$0-0
	JMP	x509_CFRelease(SB)
GLOBL	·x509_CFRelease_trampoline_addr(SB), RODATA, $8
DATA	·x509_CFRelease_trampoline_addr(SB)/8, $x509_CFRelease_trampoline<>(SB)

TEXT x509_CFDictionaryGetValueIfPresent_trampoline<>(SB),NOSPLIT,$0-0
	JMP	x509_CFDictionaryGetValueIfPresent(SB)
GLOBL	·x509_CFDictionaryGetValueIfPresent_trampoline_addr(SB), RODATA, $8
DATA	·x509_CFDictionaryGetValueIfPresent_trampoline_addr(SB)/8, $x509_CFDictionaryGetValueIfPresent_trampoline<>(SB)

TEXT x509_CFNumberGetValue_trampoline<>(SB),NOSPLIT,$0-0
	JMP	x509_CFNumberGetValue(SB)
GLOBL	·x509_CFNumberGetValue_trampoline_addr(SB), RODATA, $8
DATA	·x509_CFNumberGetValue_trampoline_addr(SB)/8, $x509_CFNumberGetValue_trampoline<>(SB)

TEXT x509_CFEqual_trampoline<>(SB),NOSPLIT,$0-0
	JMP	x509_CFEqual(SB)
GLOBL	·x509_CFEqual_trampoline_addr(SB), RODATA, $8
DATA	·x509_CFEqual_trampoline_addr(SB)/8, $x509_CFEqual_trampoline<>(SB)

TEXT x509_CFArrayCreateMutable_trampoline<>(SB),NOSPLIT,$0-0
	JMP	x509_CFArrayCreateMutable(SB)
GLOBL	·x509_CFArrayCreateMutable_trampoline_addr(SB), RODATA, $8
DATA	·x509_CFArrayCreateMutable_trampoline_addr(SB)/8, $x509_CFArrayCreateMutable_trampoline<>(SB)

TEXT x509_CFArrayAppendValue_trampoline<>(SB),NOSPLIT,$0-0
	JMP	x509_CFArrayAppendValue(SB)
GLOBL	·x509_CFArrayAppendValue_trampoline_addr(SB), RODATA, $8
DATA	·x509_CFArrayAppendValue_trampoline_addr(SB)/8, $x509_CFArrayAppendValue_trampoline<>(SB)

TEXT x509_CFDataCreate_trampoline<>(SB),NOSPLIT,$0-0
	JMP	x509_CFDataCreate(SB)
GLOBL	·x509_CFDataCreate_trampoline_addr(SB), RODATA, $8
DATA	·x509_CFDataCreate_trampoline_addr(SB)/8, $x509_CFDataCreate_trampoline<>(SB)

TEXT x509_CFErrorCopyDescription_trampoline<>(SB),NOSPLIT,$0-0
	JMP	x509_CFErrorCopyDescription(SB)
GLOBL	·x509_CFErrorCopyDescription_trampoline_addr(SB), RODATA, $8
DATA	·x509_CFErrorCopyDescription_trampoline_addr(SB)/8, $x509_CFErrorCopyDescription_trampoline<>(SB)

TEXT x509_CFErrorGetCode_trampoline<>(SB),NOSPLIT,$0-0
	JMP	x509_CFErrorGetCode(SB)
GLOBL	·x509_CFErrorGetCode_trampoline_addr(SB), RODATA, $8
DATA	·x509_CFErrorGetCode_trampoline_addr(SB)/8, $x509_CFErrorGetCode_trampoline<>(SB)

TEXT x509_CFStringCreateExternalRepresentation_trampoline<>(SB),NOSPLIT,$0-0
	JMP	x509_CFStringCreateExternalRepresentation(SB)
GLOBL	·x509_CFStringCreateExternalRepresentation_trampoline_addr(SB), RODATA, $8
DATA	·x509_CFStringCreateExternalRepresentation_trampoline_addr(SB)/8, $x509_CFStringCreateExternalRepresentation_trampoline<>(SB)
