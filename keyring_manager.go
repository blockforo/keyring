// Copyright (c) Blockforo
// SPDX-License-Identifier: BUSL-1.1

package keyring

import (
	"fmt"

	"github.com/blockforo/keyring/mlock"
)

func NewManager() KeyringManager {
	// If mlock is supported, make sure we do not swap to disk
	if mlock.Supported() {
		if err := mlock.LockAllMemory(); err != nil {
			fmt.Println("Could not call mlock. CAP_IPC_LOCK capability missing ?")
		}
	}

	// Return the
	return &keyringManager{}
}

func keyringName(name string) string {
	return fmt.Sprintf("%s:%s", KEY_TYPE_KEYRING, name)
}

// Helper method to create a new keyring with a name and link it to a parent keyring
func newKeyRing(name string, parentKeyring Serial) (Keyring, error) {

	// Otherwise create a new keyring and link it with the user
	keyringId, err := addKeyring(name, parentKeyring)

	// convert the error if necessary
	err = fromErr(err)

	// Return the serial and the error
	return &keyRing{
		name:          name,
		id:            Serial(keyringId),
		parentKeyring: parentKeyring,
	}, err
}

// Open an existing keyring via its name
func openKeyring(name string, parentKeyring Serial) (Keyring, error) {
	// Check if the keyring already exists and if so load it
	existing, err := getKeyringId(name, parentKeyring)

	// If we have no errors, then return
	if err == nil && existing > 0 {
		return &keyRing{
			name:          name,
			id:            existing,
			parentKeyring: parentKeyring,
		}, nil
	}

	// Create the new keyring
	newKeyRing, err := newKeyRing(name, parentKeyring)

	if err != nil {
		return nil, err
	}

	return newKeyRing, err

}

// // Open an existing keyring via its name
// func openKeyringWithPermissions(name string, parentKeyring Serial, permission uint32) (Keyring, error) {

// 	// Open or create the keyring
// 	newKeyRing, err := openKeyring(name, parentKeyring)

// 	// set the permissions
// 	err = setPermissions(newKeyRing.Id(), permission)
// 	if err != nil {
// 		fmt.Println("NEW KEY RING PERM ERROR", err)
// 	}

// 	return newKeyRing, err

// }

// NewUserKeyring add a new keyring with the specific name and default permissions and link it
// to the User keyring (of the user running the process)
// This is a persistent keyring (until the user is deleted)
// Default permissions: all user permissions granted and no group or third party permissions.
func (krs *keyringManager) UserKeyring(name string) (Keyring, error) {
	// permissions := buildPermissions(KEY_PERM_ALL, KEY_PERM_NONE, KEY_PERM_NONE)
	return openKeyring(keyringName(name), KEY_SPEC_USER_KEYRING)
}

// Open an existing process keyring
func (krs *keyringManager) ProcessKeyring(name string) (Keyring, error) {
	// permissions := buildPermissions(KEY_PERM_ALL, KEY_PERM_ALL, KEY_PERM_NONE)
	return openKeyring(keyringName(name), KEY_SPEC_PROCESS_KEYRING)
}

// Open an existing process session keyring
func (krs *keyringManager) SessionKeyring(name string) (Keyring, error) {
	// permissions := buildPermissions(KEY_PERM_ALL, KEY_PERM_ALL, KEY_PERM_NONE)
	return openKeyring(keyringName(name), KEY_SPEC_SESSION_KEYRING)
}

// Open the keyring for the current process user session
func (krs *keyringManager) UserSessionKeyring(name string) (Keyring, error) {
	// permissions := buildPermissions(KEY_PERM_ALL, KEY_PERM_ALL, KEY_PERM_NONE)
	return openKeyring(keyringName(name), KEY_SPEC_USER_SESSION_KEYRING)
}
