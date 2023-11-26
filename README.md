#### Status

[![Blockcert](https://github.com/blockforo/keyring/actions/workflows/build.yaml/badge.svg)](https://github.com/blockforo/keyring/actions/workflows/build.yaml)

- Needs Go 1.21

#### Keyring

A library which makes it easy to interact with the [Linux Kernel Key Retention Service](https://www.kernel.org/doc/html/v6.0/security/keys/core.html)  
It's currently being used by `blockcert` to store TLS certificate in the Linux keyring for achieving a zero-trust architecture

#### Requirements

- Unprivileged users must have the `CAP_IPC_LOCK` capability because we are preventing memory from being swapped to disk since it can contain cryptographic data
- If you want to change the permissions of a keyring you do not own you need the `CAP_SYS_ADMIN` capability (effectively `root`)

#### Keys

Keys represent units of cryptographic data, authentication tokens, keyrings.  

Each key has a number of attributes:
- A serial number: positive non-zero 32-bit integers
- A type.
- A name (for matching a key in a search).
- Access control information.
- An expiry time.
- A payload (which contains the cryptographic data).
- A State.

Each key can be in one of a number of basic states:

- **Uninstantiated**: The key exists, but does not have any data attached. Keys being requested from userspace will be in this state.
- **Instantiated**: This is the normal state. The key is fully formed, and has data attached.
- **Negative**: This is a relatively short-lived state. The key acts as a note saying that a previous call out to userspace failed, and acts as a throttle on key lookups. A negative key can be updated to a normal state.
- **Expired**: Keys can have lifetimes set. If their lifetime is exceeded, they traverse to this state. An expired key can be updated back to a normal state.
- **Revoked**: A key is put in this state by userspace action. It can’t be found or operated upon (apart from by unlinking it).
- **Dead**: The key’s type was unregistered, and so the key is now useless.

Garbage collection:
- **Dead** keys (for which the type has been removed) will be automatically unlinked from those keyrings that point to them and deleted as soon as possible by a background garbage collector.
- **Revoked** and **Expired** keys will be garbage collected, but only after a certain amount of time has passed. This time is set as a number of seconds in: `/proc/sys/kernel/keys/gc_delay`

#### Key types

- **Keyrings** are special keys that contain a list of other keys. Keyring lists can be modified using various system calls. Keyrings are not be given a payload when created.

- **User** A key of this type has a name and a payload that are arbitrary blobs of data. These can be created, updated and read by this library, and aren’t intended for use by kernel services.

- **Logon** (NOT SUPPORTED) Like a “user” key, a “logon” key has a payload that is an arbitrary blob of data. It is intended as a place to store secrets which are accessible to the kernel but not to userspace programs.

#### Linux processes

Each process subscribes to three keyrings: 
- a thread-specific keyring:  `KEY_SPEC_THREAD_KEYRING`
- a process-specific keyring: `KEY_SPEC_PROCESS_KEYRING`
- a session-specific keyring: `KEY_SPEC_SESSION_KEYRING`

The thread-specific keyring `KEY_SPEC_THREAD_KEYRING` is discarded from the child when any sort of `clone`, `fork`, `vfork` or `execve` occurs.  
A new keyring is created only when required.

The process-specific keyring `KEY_SPEC_PROCESS_KEYRING` is replaced with an empty one in the child on `clone`, `fork`, `vfork` unless `CLONE_THREAD` is supplied, in which case it is shared.  
`execve` also discards the process’s process keyring and creates a new one.

The session-specific keyring `KEY_SPEC_SESSION_KEYRING` is persistent across `clone`, `fork`, `vfork` and `execve`, even when the latter executes a `set-UID` or `set-GID` binary.  
A process can, however, replace its current session keyring with a new one by using `PR_JOIN_SESSION_KEYRING`.  
It is permitted to request an anonymous new one, or to attempt to create or join one of a specific name.

The ownership of the thread keyring changes when the real `UID` and `GID` of the thread changes.

##### Keyring types recap

- **KEY_SPEC_THREAD_KEYRING** This specifies the caller's thread-specific keyring.
- **KEY_SPEC_PROCESS_KEYRING** This specifies the caller's process-specific keyring.
- **KEY_SPEC_SESSION_KEYRING** This specifies the caller's session-specific keyring.
- **KEY_SPEC_USER_KEYRING** This specifies the caller's UID-specific keyring.
- **KEY_SPEC_USER_SESSION_KEYRING** This specifies the caller's UID-session keyring.
- **KEY_SPEC_GROUP_KEYRING** This specifies the caller's GID-specific keyring.

#### Linux users

Each user ID resident in the system holds two special keyrings: 
- a user specific keyring: `KEY_SPEC_USER_SESSION_KEYRING`
- a default user session keyring: `KEY_SPEC_USER_KEYRING` (The default session keyring is initialised with a link to the user-specific keyring.)

When a process changes its real UID, if it used to have no session key, it will be subscribed to the default session key for the new UID.

If a process attempts to access its session key when it doesn’t have one, it will be subscribed to the default for its current UID.

#### Keyring/keys permissions

Keys have an owner user ID, a group access ID, and a permissions mask.  
The mask has up to eight bits each for possessor, user, group and other access. Only six of each set of eight bits are defined.  
These permissions granted are:

- `View`: This permits a key or keyring’s attributes to be viewed - including key type and description.
- `Read`: This permits a key’s payload to be viewed or a keyring’s list of linked keys.
- `Write`: This permits a key’s payload to be instantiated or updated, or it allows a link to be added to or removed from a keyring.
- `Search`: This permits keyrings to be searched and keys to be found. Searches can only recurse into nested keyrings that have search permission set.
- `Link`: This permits a key or keyring to be linked to. To create a link from a keyring to a key, a process must have Write permission on the keyring and Link permission on the key.
- `Set Attribute`: This permits a key’s UID, GID and permissions mask to be changed.

#### Keys quotas

Each user has two quotas against which the keys they own are tracked:  
- One limits the total number of keys and keyrings
- the other limits the total amount of description and payload space that can be consumed.

Four new sysctl files have been added also for the purpose of controlling the quota limits on keys:

- `/proc/sys/kernel/keys/root_maxkeys`: This file hold the maximum number of keys that root may have
- `/proc/sys/kernel/keys/root_maxbytes`: This file hold the maximum total number of bytes of data that root may have stored in those keys.
- `/proc/sys/kernel/keys/maxkeys`: This file hold the maximum number of keys that each non-root user may have
- `/proc/sys/kernel/keys/maxbytes`: This file hold the maximum total number of bytes of data that each of those users may have stored in their keys.

`root` may alter these by writing each new limit as a decimal number string to the appropriate file.
