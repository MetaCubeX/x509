// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build darwin && arm64

#include "textflag.h"

// syscall.syscall6 passes all arguments in integer registers. Move the bit
// representation of CFAbsoluteTime from the second integer argument to F0.
TEXT x509_CFDateCreate_trampoline<>(SB),NOSPLIT,$0-0
	FMOVD	R1, F0
	JMP	x509_CFDateCreate(SB)
GLOBL	·x509_CFDateCreate_trampoline_addr(SB), RODATA, $8
DATA	·x509_CFDateCreate_trampoline_addr(SB)/8, $x509_CFDateCreate_trampoline<>(SB)
