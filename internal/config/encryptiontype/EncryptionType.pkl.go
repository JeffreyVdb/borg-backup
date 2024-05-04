// Code generated from Pkl module `co.vandenborne.BackupConfig`. DO NOT EDIT.
package encryptiontype

import (
	"encoding"
	"fmt"
)

type EncryptionType string

const (
	None                EncryptionType = "none"
	Keyfile             EncryptionType = "keyfile"
	Repokey             EncryptionType = "repokey"
	KeyfileBlake2       EncryptionType = "keyfile-blake2"
	RepokeyBlake2       EncryptionType = "repokey-blake2"
	AuthenticatedBlake2 EncryptionType = "authenticated-blake2"
)

// String returns the string representation of EncryptionType
func (rcv EncryptionType) String() string {
	return string(rcv)
}

var _ encoding.BinaryUnmarshaler = new(EncryptionType)

// UnmarshalBinary implements encoding.BinaryUnmarshaler for EncryptionType.
func (rcv *EncryptionType) UnmarshalBinary(data []byte) error {
	switch str := string(data); str {
	case "none":
		*rcv = None
	case "keyfile":
		*rcv = Keyfile
	case "repokey":
		*rcv = Repokey
	case "keyfile-blake2":
		*rcv = KeyfileBlake2
	case "repokey-blake2":
		*rcv = RepokeyBlake2
	case "authenticated-blake2":
		*rcv = AuthenticatedBlake2
	default:
		return fmt.Errorf(`illegal: "%s" is not a valid EncryptionType`, str)
	}
	return nil
}
