// Copyright (c) Blockforo
// SPDX-License-Identifier: BUSL-1.1

package keyring

// buildPermissions constructs the permission mask from the parameters.
func buildPermissions(possessor, user, group, others uint32) KeyPerm {
	perm := others << keyPermissionOthers
	perm |= user << keyPermissionUser
	perm |= group << keyPermissionGroup
	perm |= possessor << keyPermissionPosessor

	return KeyPerm(perm)
}
