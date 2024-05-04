package borg

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/JeffreyVdb/borg-backup/pkg/borg/archive"
	"os"
	"os/exec"
	"strings"
)

type Config struct {
	Executable string
	Encryption string
}

type Borg struct {
	Executable string
	Encryption string
	version    string
}

func New(cfg Config) (*Borg, error) {
	if cfg.Executable == "" {
		executablePath, err := exec.LookPath("borg")
		if err != nil {
			return nil, fmt.Errorf("executable wasn't given and couldn't find borg in PATH: %w", err)
		}

		cfg.Executable = executablePath
	}

	if cfg.Encryption == "" {
		cfg.Encryption = "none"
	}

	version, err := determineBorgVersion(cfg.Executable)
	if err != nil {
		return nil, fmt.Errorf("failed to determine borg version: %w", err)
	}

	return &Borg{
		Executable: cfg.Executable,
		Encryption: cfg.Encryption,
		version:    version,
	}, nil
}

func (b Borg) IsV1() bool {
	return strings.HasPrefix(b.version, "1.")
}

func (b Borg) IsV2() bool {
	return strings.HasPrefix(b.version, "2.")
}

func (b Borg) IsRepository(repoPath string) (bool, error) {
	if b.IsV1() {
		cmd := exec.Command(b.Executable, "info", repoPath)
		err := cmd.Run()
		if err != nil {
			if exitErr := (&exec.ExitError{}); errors.As(err, &exitErr) {
				return false, nil
			}

			return false, fmt.Errorf("failed to check if path is a borg repository: %w", err)
		}
	} else if b.IsV2() {
		return false, fmt.Errorf("borg v2 not supported yet")
	} else {
		return false, fmt.Errorf("unsupported borg version: %s", b.version)
	}

	return true, nil

}

func (b Borg) CreateRepository(repoPath string) error {
	if b.IsV1() {
		cmd := exec.Command(b.Executable, "init", "--encryption", b.Encryption, repoPath)
		out, err := cmd.CombinedOutput()
		if err != nil {
			return fmt.Errorf("failed to create borg repository: %s: %w", out, err)
		}
	} else if b.IsV2() {
		return fmt.Errorf("borg v2 not supported yet")
	} else {
		return fmt.Errorf("unsupported borg version: %s", b.version)
	}

	return nil
}

func (b Borg) CreateArchive(repoPath, archiveName string, backupPaths []string, archiveOptions ...archive.OptionFunc) error {
	options := archive.Options{}
	for _, option := range archiveOptions {
		option(&options)
	}

	if options.Compression == "" {
		options.Compression = "none"
	}

	cmd := exec.Command(b.Executable, "create", "--compression", options.Compression, fmt.Sprintf("%s::%s", repoPath, archiveName))
	if options.WorkingDirectory != "" {
		cmd.Dir = options.WorkingDirectory
	}

	cmd.Args = append(cmd.Args, backupPaths...)
	for _, exclude := range options.Excludes {
		cmd.Args = append(cmd.Args, "--exclude="+exclude)
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to create borg archive: %w", err)
	}

	return nil
}

func determineBorgVersion(executable string) (string, error) {
	cmd := exec.Command(executable, "--version")
	out, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("failed to determine borg version: %w", err)
	}

	fields := bytes.Fields(out)
	if len(fields) < 2 {
		return "", fmt.Errorf("unexpected output from borg --version: %s", out)
	}

	return string(fields[1]), nil
}
