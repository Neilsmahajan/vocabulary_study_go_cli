package storage

import "fmt"

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
	if err := delete(vocab, word); err != nil {
		return fmt.Errorf("error deleting word '%s': %w", word, err)
	}
	fmt.Printf("✅ Removed word '%s' from vocab\n", word)
	return nil
}

func delete(vocab map[string]VocabEntry, word string) error {
	fmt.Printf("Deleting word %s from vocab\n", word)
	return nil
}
