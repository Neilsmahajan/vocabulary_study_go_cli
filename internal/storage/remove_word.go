package storage

import (
	"encoding/json"
	"fmt"
	"os"
)

func RemoveWord(path, word string) error {
	if word == "" {
		return fmt.Errorf("word is required")
	}
	vocab, err := LoadVocab(path)
	if err != nil {
		return fmt.Errorf("error loading vocab: %w", err)
	}
	if _, exists := vocab[word]; !exists {
		return fmt.Errorf("word '%s' does not exist in vocab", word)
	}
	if err := delete(vocab, word, path); err != nil {
		return fmt.Errorf("error deleting word '%s': %w", word, err)
	}
	fmt.Printf("✅ Removed word '%s' from vocab\n", word)
	return nil
}

func delete(vocab map[string]VocabEntry, word, path string) error {
	updatedVocab := make(map[string]VocabEntry)
	for k, v := range vocab {
		if k != word {
			updatedVocab[k] = v
		}
	}
	if len(updatedVocab) == len(vocab) {
		return fmt.Errorf("word '%s' not found in vocab", word)
	}
	vocab = updatedVocab
	updatedData, err := json.MarshalIndent(vocab, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshaling vocab: %w", err)
	}
	if err := os.WriteFile(path, updatedData, 0644); err != nil {
		return fmt.Errorf("error writing updated vocab to file: %w", err)
	}

	fmt.Printf("Deleting word %s from vocab\n", word)
	return nil
}
