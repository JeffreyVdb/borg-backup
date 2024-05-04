// Code generated from Pkl module `co.vandenborne.BackupConfig`. DO NOT EDIT.
package config

import "github.com/JeffreyVdb/borg-backup/internal/config/encryptiontype"

type Repository struct {
	Path string `pkl:"path"`

	Encryption encryptiontype.EncryptionType `pkl:"encryption"`
}
