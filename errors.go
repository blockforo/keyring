// Copyright (c) 2023 Blockforo
// SPDX-License-Identifier: BUSL-1.1

package keyring

import (
	"errors"

	"golang.org/x/sys/unix"
)

var (
	// Key errors
	// --------------------------------------------------------------------------------------------------

	// ErrKeyNotFound ENOENT the specific key was not found
	ErrKeyNotFound = errors.New("key not found")

	// ErrKeyNotPresent ENOKEY in case of key deletion and the key is not present
	ErrKeyNotPresent = errors.New("key not present")

	// ErrKeyUpdate EOPNOTSUPP when a key update happens and the specific key type does not suppor the update
	// operation
	ErrKeyUpdate = errors.New("key type does not support update operation")

	// ErrKeyPermission EINVAL when an operation on a key fails because of its permissions
	ErrKeyPermission = errors.New("invalid key permissions")

	// ErrKeyAccessError EACCES when the key does not have access permissions
	ErrKeyAccessError = errors.New("access permission missing")

	// ErrKeyAlreadyExists EEXIST when a key already exists
	ErrKeyAlreadyExists = errors.New("key exists")

	// ErrKeyRevoked EKEYREVOKED when a key is revoked
	ErrKeyRevoked = errors.New("key revoked")

	// ErrKeyExpired EKEYEXPIRED when a key is expired
	ErrKeyExpired = errors.New("key expired")

	// Generic permission error
	ErrKeyDeletionError = errors.New("key deletion error")

	// Keyring errors
	// --------------------------------------------------------------------------------------------------

	// ErrInvalidKeyring ENOTDIR when an operatrion is executed on a keyring and the keyring does not exist
	ErrInvalidKeyring = errors.New("key is not keyring")

	// ErrKeyringIsFull ENFILE when the keyring is full (driven by cgroup)
	ErrKeyringIsFull = errors.New("keyring full")

	// ErrKeyringLinkTooDeep ELOOP when linking a keyring to another keyring which is too deep
	ErrKeyringLinkTooDeep = errors.New("keyring link too deep")

	// ErrKeyringLinkCycle EDEADLK when linking keyrings causes a cycle
	ErrKeyringLinkCycle = errors.New("keyring link cycle")
)

// Converts the error to the internal one, for an easier handling
func fromErr(err error) error {

	// Short circuit
	if err == nil {
		return nil
	}

	if errors.Is(err, unix.ENOENT) {
		return ErrKeyNotFound
	}

	if errors.Is(err, unix.ENOKEY) {
		return ErrKeyNotPresent
	}

	if errors.Is(err, unix.EOPNOTSUPP) {
		return ErrKeyUpdate
	}

	if errors.Is(err, unix.EINVAL) {
		return ErrKeyPermission
	}

	if errors.Is(err, unix.EACCES) {
		return ErrKeyAccessError
	}

	if errors.Is(err, unix.EEXIST) {
		return ErrKeyAlreadyExists
	}

	if errors.Is(err, unix.EKEYREVOKED) {
		return ErrKeyRevoked
	}

	if errors.Is(err, unix.EKEYEXPIRED) {
		return ErrKeyExpired
	}

	if errors.Is(err, unix.ENOTDIR) {
		return ErrInvalidKeyring
	}

	if errors.Is(err, unix.ENFILE) {
		return ErrKeyringIsFull
	}

	if errors.Is(err, unix.ELOOP) {
		return ErrKeyringLinkTooDeep
	}

	if errors.Is(err, unix.EDEADLK) {
		return ErrKeyringLinkCycle
	}

	// Otherwise just return the original error
	return err
}
