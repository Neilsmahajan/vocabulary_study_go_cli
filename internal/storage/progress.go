package storage

import (
	"encoding/json"
	"fmt"
	"os"
)

func LoadProgress(path string) (map[string]string, error) {
	progress := make(map[string]string)

	// Check if file exists
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return progress, nil
	}

	// Read the file
	bytes, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read progress file: %w", err)
	}

	if err := json.Unmarshal(bytes, &progress); err != nil {
		return nil, fmt.Errorf("failed to parse progress JSON: %w", err)
	}
	return progress, nil
}

func SaveProgress(path string, progress map[string]string) error {
	bytes, err := json.MarshalIndent(progress, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to encode progress to JSON: %w", err)
	}

	if err := os.WriteFile(path, bytes, 0644); err != nil {
		return fmt.Errorf("failed to write progress file: %w", err)
	}
	return nil
}
