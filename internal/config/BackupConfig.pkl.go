// Code generated from Pkl module `co.vandenborne.BackupConfig`. DO NOT EDIT.
package config

import (
	"context"

	"github.com/apple/pkl-go/pkl"
)

type BackupConfig struct {
	Repository *Repository `pkl:"repository"`

	Cwd *string `pkl:"cwd"`

	ArchiveName string `pkl:"archiveName"`

	Compression string `pkl:"compression"`

	Paths []string `pkl:"paths"`

	Excludes []string `pkl:"excludes"`
}

// LoadFromPath loads the pkl module at the given path and evaluates it into a BackupConfig
func LoadFromPath(ctx context.Context, path string) (ret *BackupConfig, err error) {
	evaluator, err := pkl.NewEvaluator(ctx, pkl.PreconfiguredOptions)
	if err != nil {
		return nil, err
	}
	defer func() {
		cerr := evaluator.Close()
		if err == nil {
			err = cerr
		}
	}()
	ret, err = Load(ctx, evaluator, pkl.FileSource(path))
	return ret, err
}

// Load loads the pkl module at the given source and evaluates it with the given evaluator into a BackupConfig
func Load(ctx context.Context, evaluator pkl.Evaluator, source *pkl.ModuleSource) (*BackupConfig, error) {
	var ret BackupConfig
	if err := evaluator.EvaluateModule(ctx, source, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
