// Copyright (c) Blockforo
// SPDX-License-Identifier: BUSL-1.1

//go:build dragonfly || freebsd || linux || openbsd || solaris

package mlock

import (
	"syscall"

	"golang.org/x/sys/unix"
)

// It's supported for these platforms:
// dragonfly || freebsd || linux || openbsd || solaris
func init() {
	supported = true
}

// lockAllMemory locks all the process memory
func lockAllMemory() error {
	return unix.Mlockall(syscall.MCL_CURRENT | syscall.MCL_FUTURE)
}

// lockData locks only specific data
func lockData(data []byte) error {
	return unix.Mlock(data)
}

// unlockData unlocks specific data
func unlockData(data []byte) error {
	return unix.Munlock(data)
}
