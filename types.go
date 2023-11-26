// Copyright (c) Blockforo
// SPDX-License-Identifier: BUSL-1.1

package keyring

import "fmt"

// This is the unique ID of the keys (key and keyring)
type Serial int

// Int return the int representation
func (s Serial) Int() int {
	return int(s)
}

// The type of key (user, keyring)
type KeyType int

// Internal struct for the keyring manager
type keyringManager struct{}

type KeyMetadata struct {
	// The key type
	Type string

	// The key name
	Name string

	// The user owner of the key
	Uid int

	// The group owner of the key
	Gid int

	Permissions KeyPerm
}

func (m KeyMetadata) String() string {
	return fmt.Sprintf("type: %s | name: %s | uid: %d | gid: %d | perm: %s", m.Type, m.Name, m.Uid, m.Gid, m.Permissions)
}
