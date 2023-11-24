package op_stack

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

var DevnetPresetId = 901

func DevnetPreset() (*Preset, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	root, err := findMonorepoRoot(cwd)
	if err != nil {
		return nil, err
	}

	devnetFilepath := filepath.Join(root, ".devnet", "addresses.json")
	if _, err := os.Stat(devnetFilepath); errors.Is(err, fs.ErrNotExist) {
		return nil, err
	}

	content, err := os.ReadFile(devnetFilepath)
	if err != nil {
		return nil, err
	}

	var l1Contracts L1Contracts
	if err := json.Unmarshal(content, &l1Contracts); err != nil {
		return nil, err
	}

	return &Preset{
		Name:        "Local Devnet",
		OpContracts: OpContracts{Preset: DevnetPresetId, L1Contracts: l1Contracts},
	}, nil
}

func findMonorepoRoot(startDir string) (string, error) {
	dir, err := filepath.Abs(startDir)
	if err != nil {
		return "", err
	}
	for {
		modulePath := filepath.Join(dir, "go.mod")
		if _, err := os.Stat(modulePath); err == nil {
			return dir, nil
		}
		parentDir := filepath.Dir(dir)
		if parentDir == dir {
			break
		}
		dir = parentDir
	}
	return "", fmt.Errorf("monorepo root not found")
}
