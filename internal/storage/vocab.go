package storage

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type VocabEntry struct {
	PartOfSpeech    string `json:"part_of_speech"`
	Definition      string `json:"definition"`
	ExampleSentence string `json:"example_sentence"`
}

func LoadVocab(path string) (map[string]VocabEntry, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open vocab file: %w", err)
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("failed to read vocab file %w", err)
	}
	var vocab map[string]VocabEntry
	if err := json.Unmarshal(bytes, &vocab); err != nil {
		return nil, fmt.Errorf("failed to parse vocab json: %w", err)
	}
	return vocab, nil
}
