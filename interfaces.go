// Copyright (c) Blockforo
// SPDX-License-Identifier: BUSL-1.1

package keyring

type KeyringManager interface {

	// Load or create a keyring with parent the User keyring
	UserKeyring(name string) (Keyring, error)

	// Load or create a keyring with parent the process session keyring
	SessionKeyring(name string) (Keyring, error)

	// Open an existing process keyring
	ProcessKeyring(name string) (Keyring, error)

	// Open the keyring for the current process user session
	UserSessionKeyring(name string) (Keyring, error)
}

// The keyring representation
type Keyring interface {

	// Returns the name of the keyring
	Name() string

	// Returns the ID of the keyring
	Id() Serial

	// Returns the ID of the parent keyring
	Parent() Serial

	// Adds or updates a key to the key ring and returns its ID
	SetKey(name string, data []byte) (Serial, error)

	// Sets specific key permissions
	SetKeyPermissions(name string, user, group, others KeyPerm) error

	// Set the keyring permissions
	SetPermissions(name string, user, group, others KeyPerm) error

	// Gets a key from the key ring, via its name
	GetKey(name string) (Key, error)

	// Checks whether the key ring has a specific key
	HasKey(name string) bool

	// Deletes the specific key
	DeleteKey(name string) error

	// Destroys the key ring itself
	Destroy() error

	// Describe the keyring
	Describe() (KeyMetadata, error)

	// Describe the key
	DescribeKey(name string) (KeyMetadata, error)
}

// The key representation
type Key interface {
	// Returns the content of the key
	Payload() []byte

	// Returns the name of the key
	Name() string

	// Returns the ID of the key
	Id() Serial
}
