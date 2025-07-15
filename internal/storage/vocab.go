package storage

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type VocabEntry struct {
	PartOfSpeech    string `json:"part_of_speech"`
	Definition      string `json:"definition"`
	ExampleSentence string `json:"sentence"`
}

func LoadVocab(path string) (map[string]VocabEntry, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open vocab file: %w", err)
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("Failed to read vocab file %w", err)
	}
	var vocab map[string]VocabEntry
	if err := json.Unmarshal(bytes, &vocab); err != nil {
		return nil, fmt.Errorf("Failed to parse vocab JSON: %w", err)
	}
	return vocab, nil
}
