// Copyright (c) Blockforo
// SPDX-License-Identifier: BUSL-1.1

package keyring

import (
	"fmt"
)

// Internal struct to hold the key ring
type keyRing struct {

	// The unique id of the keyring
	id Serial

	// The name of the keyring
	name string

	// The parent keyring it is attached to
	parentKeyring Serial
}

func (kr *keyRing) keyName(name string) string {
	return fmt.Sprintf("%s:%s", KEY_TYPE_USER, name)
}

// Returns the name of the keyring
func (kr *keyRing) Name() string {
	return kr.name
}

// Returns the ID of the keyring
func (kr *keyRing) Id() Serial {
	return kr.id
}

// Returns the ID of the parent keyring
func (kr *keyRing) Parent() Serial {
	return kr.parentKeyring
}

// Sets the permissions of a key
func (kr *keyRing) SetKeyPermissions(name string, user, group, others KeyPerm) error {
	permissions := buildPermissions(PermAll.Uint32(), user.Uint32(), group.Uint32(), others.Uint32())

	if k, err := kr.GetKey(name); err != nil {
		return err
	} else {
		return setPermissions(k.Id(), permissions.Uint32())
	}

}

// Set the permissions of the keyring
func (kr *keyRing) SetPermissions(name string, user, group, others KeyPerm) error {
	permissions := buildPermissions(PermAll.Uint32(), user.Uint32(), group.Uint32(), others.Uint32())
	return setPermissions(kr.Id(), permissions.Uint32())
}

// Returns the ID of the parent keyring
func (kr *keyRing) SetKey(name string, data []byte) (Serial, error) {
	// add the key
	id, err := addKey(kr.keyName(name), data, kr.id)

	// return the id and the error
	return Serial(id), err
}

func (kr *keyRing) Describe() (KeyMetadata, error) {
	return describe(kr.id)
}

func (kr *keyRing) DescribeKey(name string) (KeyMetadata, error) {
	existing, err := kr.GetKey(name)
	if err != nil {
		return KeyMetadata{}, err
	}

	return describe(existing.Id())
}

// Checks whether the key exists
func (kr *keyRing) HasKey(name string) bool {
	// Get its serial
	serial, err := getKeyId(kr.keyName(name), kr.id)

	// Make sure it exists
	return err == nil && serial.Int() > 0
}

// Returns the ID of the parent keyring
func (kr *keyRing) GetKey(name string) (Key, error) {

	// Get its serial
	serial, err := getKeyId(kr.keyName(name), kr.id)

	// handle the error
	if err != nil {
		return nil, err
	}

	// Get the key
	storedKey, err := getKey(Serial(serial.Int()))
	if err != nil {
		return nil, err
	}

	return &key{
		id:      serial,
		name:    name,
		payload: storedKey,
	}, err
}

// Deletes the specific key
func (kr *keyRing) DeleteKey(name string) error {
	return deleteKeyWithName(kr.keyName(name), kr.id)
}

// Destroys the key ring itself
func (kr *keyRing) Destroy() error {
	return deleteKeyring(kr.id, kr.parentKeyring)
}
