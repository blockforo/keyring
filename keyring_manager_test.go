// Copyright (c) Blockforo
// SPDX-License-Identifier: BUSL-1.1

package keyring

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKeyringAddWithPermissions(t *testing.T) {

	// Test data
	keyringName := "user_ring"
	keyName := "my_password"
	secret := "a quick brown fox jumps over the lazy dog"

	// New keyring manager
	service := NewManager()

	// New user keyring
	kr, err := service.UserKeyring(keyringName)
	assert.Nil(t, err)

	// Create a key
	id, err := kr.SetKey(keyName, []byte(secret))
	if err != nil {
		fmt.Println(err)
	}
	assert.Nil(t, err)
	assert.True(t, id > 0)

	metadata, err := kr.Describe()
	assert.Nil(t, err)
	assert.NotEmpty(t, metadata.Permissions.Uint32())
	assert.NotEmpty(t, metadata.Name)
	assert.Equal(t, KEY_TYPE_KEYRING, metadata.Type)
	assert.NotEmpty(t, metadata.Uid)
	assert.NotEmpty(t, metadata.Gid)

	// Set the key permissions
	err = kr.SetKeyPermissions(keyName, PermAll, PermRead|PermSearch, PermNone)
	assert.Nil(t, err)

	// Make sure it's present
	hasKey := kr.HasKey(keyName)
	assert.True(t, hasKey)

	// Describe the key
	metadata, err = kr.DescribeKey(keyName)
	assert.Nil(t, err)
	assert.Nil(t, err)
	assert.NotEmpty(t, metadata.Permissions.Uint32())
	assert.NotEmpty(t, metadata.Name)
	assert.Equal(t, KEY_TYPE_USER, metadata.Type)
	assert.NotEmpty(t, metadata.Uid)
	assert.NotEmpty(t, metadata.Gid)

	expectedPermissions := "3f3f0a00"
	expected := new(big.Int)
	expected.SetString(expectedPermissions, 16)
	assert.Equal(t, uint32(expected.Uint64()), metadata.Permissions.Uint32())

	// Retrieve the key
	storedKey, err := kr.GetKey(keyName)
	assert.Nil(t, err)
	assert.Equal(t, secret, string(storedKey.Payload()))

	// Delete the key
	err = kr.DeleteKey(keyName)
	if err != nil {
		fmt.Println(err)
	}
	assert.Nil(t, err)

	// Destroy the key ring
	err = kr.Destroy()
	if err != nil {
		fmt.Println(err)
	}
	assert.Nil(t, err)
}
