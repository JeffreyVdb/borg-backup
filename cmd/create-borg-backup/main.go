package main

import (
	"context"
	"fmt"
	"github.com/JeffreyVdb/borg-backup/internal/config"
	"github.com/JeffreyVdb/borg-backup/pkg/borg"
	"github.com/JeffreyVdb/borg-backup/pkg/borg/archive"
	"log"
	"os"
)

func main() {
	ctx := context.Background()
	cfg, err := loadConfig(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	_ = cfg

	borgClient, err := borg.New(borg.Config{Encryption: cfg.Repository.Encryption.String()})
	if err != nil {
		log.Fatalln(err)
	}

	isRepo, err := borgClient.IsRepository(cfg.Repository.Path)
	if err != nil {
		log.Fatalln(err)
	}

	if !isRepo {
		if err = borgClient.CreateRepository(cfg.Repository.Path); err != nil {
			log.Fatalln(err)
		}
	}

	opts := []archive.OptionFunc{
		archive.WithCompression(cfg.Compression),
		archive.WithExcludes(cfg.Excludes),
	}
	if cfg.Cwd != nil {
		opts = append(opts, archive.WithWorkingDirectory(*cfg.Cwd))
	}
	if err = borgClient.CreateArchive(cfg.Repository.Path, cfg.ArchiveName, cfg.Paths, opts...); err != nil {
		log.Fatalln(err)
	}
}

func loadConfig(ctx context.Context) (*config.BackupConfig, error) {
	if len(os.Args) != 2 {
		return nil, fmt.Errorf("usage: %s <config.pkl>", os.Args[0])
	}

	cfg, err := config.LoadFromPath(ctx, os.Args[1])
	if err != nil {
		return nil, fmt.Errorf("failed to load config file %s: %w", os.Args[1], err)
	}

	return cfg, nil
}
