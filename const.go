// Copyright (c) Blockforo
// SPDX-License-Identifier: BUSL-1.1

package keyring

import (
	"fmt"
	"math/bits"

	"golang.org/x/sys/unix"
)

// The ker permissions
type KeyPerm uint32

// The entity to which the permissions are applied:
// user, group, others
type entityPerm uint32

//nolint:revive
const (
	PermNone       KeyPerm = 0            // -: no permissions
	PermView       KeyPerm = 1 << 0       // v: view a key or keyring’s attributes - including key type and description.
	PermRead       KeyPerm = 1 << 1       // r: read the key payload or a keyring’s list of linked keys.
	PermWrite      KeyPerm = 1 << 2       // w: write key’s payload, or add and remove links to or from a keyring
	PermSearch     KeyPerm = 1 << 3       // s: search keyrings and keys. Recurse into nested keyrings.
	PermLink       KeyPerm = 1 << 4       // l: link keys or keyrings. To create a link from a keyring to a key, a process must have Write permission on the keyring and Link permission on the key.
	PermAttributes KeyPerm = 1 << 5       // a: allows to change a key’s UID, GID and permissions mask
	PermAll        KeyPerm = (1 << 6) - 1 // --alswrv: all the permissions

	// KEY_PERM_VIEW    = uint32(1 << 0)
	// KEY_PERM_READ    = uint32(1 << 1)
	// KEY_PERM_WRITE   = uint32(1 << 2)
	// KEY_PERM_SEARCH  = uint32(1 << 3)
	// KEY_PERM_LINK    = uint32(1 << 4)
	// KEY_PERM_SETATTR = uint32(1 << 5)
	// KEY_PERM_ALL     = uint32((1 << 6) - 1)

	keyPermissionPosessor entityPerm = 24
	keyPermissionUser     entityPerm = 16
	keyPermissionGroup    entityPerm = 8
	keyPermissionOthers   entityPerm = 0

	KEYCTL_PERM_OTHERS  = 0
	KEYCTL_PERM_GROUP   = 8
	KEYCTL_PERM_USER    = 16
	KEYCTL_PERM_PROCESS = 24

	// thread-specific keyring: -1
	KEY_SPEC_THREAD_KEYRING = unix.KEY_SPEC_THREAD_KEYRING

	// process-specific keyring: -2
	KEY_SPEC_PROCESS_KEYRING = unix.KEY_SPEC_PROCESS_KEYRING

	// session-specific keyring: -3
	KEY_SPEC_SESSION_KEYRING = unix.KEY_SPEC_SESSION_KEYRING

	// UID-specific keyring: -4
	KEY_SPEC_USER_KEYRING = unix.KEY_SPEC_USER_KEYRING

	// UID-session keyring: -5
	KEY_SPEC_USER_SESSION_KEYRING = unix.KEY_SPEC_USER_SESSION_KEYRING

	// GID-specific keyring: -6
	KEY_SPEC_GROUP_KEYRING = unix.KEY_SPEC_GROUP_KEYRING

	// KEY_TYPE_USER
	// A key of this type has a description and a payload that are arbitrary blobs of data.
	// These can be created, updated and read by userspace, and aren’t intended for use by kernel services.
	KEY_TYPE_USER = "user"

	// KEY_TYPE_KEYRING
	// Keyrings are special keys that contain a list of other keys.
	// Keyring lists can be modified using various system calls.
	// Keyrings should not be given a payload when created.
	KEY_TYPE_KEYRING = "keyring"
)

func (p KeyPerm) String() string {

	base := p.Uint32()

	posessor := bits.Reverse8(byte((base >> 24)))
	prosessorString := processByte(posessor)

	user := bits.Reverse8(byte((base >> 16)))
	userString := processByte(user)

	group := bits.Reverse8(byte((base >> 8)))
	groupString := processByte(group)

	others := bits.Reverse8(byte((base >> 0)))
	othersString := processByte(others)

	return fmt.Sprintf("owner[%s] | user[%s] | group[%s] | others[%s]", prosessorString, userString, groupString, othersString)
}

func processByte(data byte) string {
	// The first 2 bytes are not used bby the kernel
	const str = "--alswrv"

	w := 0
	var buf [8]byte
	for i, c := range str {
		result := data & (1 << (uint(i) & 0xff))
		if result != 0 {
			buf[w] = byte(c)
		} else {
			buf[w] = '-'
		}
		w++
	}
	return string(buf[:w])
}

func (p KeyPerm) Uint32() uint32 {
	return uint32(p)
}
