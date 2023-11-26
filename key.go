// Copyright (c) Blockforo
// SPDX-License-Identifier: BUSL-1.1

package keyring

// Internal keyring structure
type key struct {

	// The keyring unique id
	id Serial

	// The name of the keyring
	name string

	// The cryptographic material, if this is a key, empty otherwise
	payload []byte
}

// =================================================================
// Key interface
// Returns the ID of the key
func (k *key) Id() Serial {
	return k.id
}

// Returns the name of the key
func (k *key) Name() string {
	return k.name
}

// Returns the content of the key
func (k *key) Payload() []byte {
	return k.payload
}
