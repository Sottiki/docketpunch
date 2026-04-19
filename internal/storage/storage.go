package storage

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/Sottiki/docketpunch/internal/docket"
)

func getDocketPath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(homeDir, ".docket", "tasks.json"), nil
}

func Init() error {
	path, err := getDocketPath()
	if err != nil {
		return err
	}

	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return Save(docket.NewDocket())
	}

	return nil
}

func Save(d *docket.Docket) error {
	path, err := getDocketPath()
	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}

func Load() (*docket.Docket, error) {
	if err := Init(); err != nil {
		return nil, err
	}

	path, err := getDocketPath()
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var d docket.Docket
	if err := json.Unmarshal(data, &d); err != nil {
		return nil, err
	}

	return &d, nil
}
