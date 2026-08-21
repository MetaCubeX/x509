// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build darwin

package macos

import (
	"syscall"
	_ "unsafe"
)

// syscall_syscall6 is implemented by package syscall using runtime.libcCall.
// This matches the bridge used by golang.org/x/sys/unix on Darwin.
func syscall_syscall6(fn, a1, a2, a3, a4, a5, a6 uintptr) (r1, r2 uintptr, err syscall.Errno)

//go:linkname syscall_syscall6 syscall.syscall6
