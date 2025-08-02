package storage

import (
	"encoding/json"
	"fmt"
	"os"
)

func AddWord(path, word, pos, definition, example string) error {
	if word == "" || pos == "" || definition == "" || example == "" {
		return fmt.Errorf("all fields (word, pos, definition, example) are required")
	}
	vocab, err := LoadVocab(path)
	if err != nil {
		return fmt.Errorf("error loading vocab: %w", err)
	}
	if _, exists := vocab[word]; exists {
		return fmt.Errorf("word '%s' already exists in vocab", word)
	}
	vocab[word] = VocabEntry{
		PartOfSpeech:    pos,
		Definition:      definition,
		ExampleSentence: example,
	}
	if err := saveVocab(path, word, vocab[word]); err != nil {
		return fmt.Errorf("error saving vocab: %w", err)
	}
	fmt.Printf("✅ Added word '%s' to vocab\n", word)
	return nil
}

func saveVocab(path, word string, entry VocabEntry) error {
	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("failed to open vocab file: %w", err)
	}
	defer file.Close()
	// Add logic to write the entry to the json path file
	fileData, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("failed to read vocab file: %w", err)
	}
	var vocab map[string]VocabEntry
	if err = json.Unmarshal(fileData, &vocab); err != nil {
		return fmt.Errorf("failed to unmarshal vocab data: %w", err)
	}
	vocab[word] = entry
	fileData, err = json.MarshalIndent(vocab, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal vocab data: %w", err)
	}
	if err := os.WriteFile(path, fileData, 0644); err != nil {
		return fmt.Errorf("failed to write vocab data to file: %w", err)
	}
	fmt.Printf("Saving vocab word %s to %s: %+v\n", word, path, entry)
	return nil
}
