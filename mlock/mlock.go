// Copyright (c) Blockforo
// SPDX-License-Identifier: BUSL-1.1

package mlock

// This is a local variable which is assigned for each platform that
// support mlock
var supported bool

// Returns whether the current platform supports mlock
func Supported() bool {
	return supported
}

// LockAllMemory locks all the process memory
func LockAllMemory() error {
	return lockAllMemory()
}

// LockData locks only specific data
func LockData(data []byte) error {
	return lockData(data)
}

// UnLockData unlocks specific data
func UnLockData(data []byte) error {
	return unlockData(data)
}
