// Copyright (c) Blockforo
// SPDX-License-Identifier: BUSL-1.1

//go:build darwin || nacl || netbsd || plan9 || windows || js

package mlock

// It's unsupported for these platforms:
// darwin || nacl || netbsd || plan9 || windows || js
func init() {
	supported = false
}

// lockAllMemory locks all the process memory
func lockAllMemory() error {
	// Mlockall prevents all current and future pages from being swapped to the disk
	return nil
}

// lockData locks only specific data
func lockData(data []byte) error {
	return nil
}

// unlockData unlocks specific data
func unlockData(data []byte) error {
	return nil
}
