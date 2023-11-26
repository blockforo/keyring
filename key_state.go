// Copyright (c) Blockforo
// SPDX-License-Identifier: BUSL-1.1

package keyring

type KeyState int

const (

	// Uninstantiated. The key exists, but does not have any data attached.
	// Keys being requested from userspace will be in this state.
	Uninstantiated KeyState = iota

	// Instantiated. This is the normal state. The key is fully formed, and has data attached.
	Instantiated

	// Negative. This is a relatively short-lived state.
	// The key acts as a note saying that a previous call out to userspace failed,
	// and acts as a throttle on key lookups.
	// A negative key can be updated to a normal state.
	Negative

	// Expired. Keys can have lifetimes set.
	// If their lifetime is exceeded, they traverse to this state.
	// An expired key can be updated back to a normal state.
	// Garbage collected
	Expired

	// Revoked. A key is put in this state by userspace action.
	// It can’t be found or operated upon (apart from by unlinking it).
	// Garbage collected
	Revoked

	// Dead. The key’s type was unregistered, and so the key is now useless.
	// Garbage collected
	Dead
)
