package storage

import "fmt"

func AddWord(vocabPath, word, pos, definition, example string) error {
	if word == "" || pos == "" || definition == "" || example == "" {
		return fmt.Errorf("all fields (word, pos, definition, example) are required")
	}
	vocab, err := LoadVocab(vocabPath)
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
	if err := saveVocab(vocabPath, word, vocab[word]); err != nil {
		return fmt.Errorf("error saving vocab: %w", err)
	}
	fmt.Printf("✅ Added word '%s' to vocab\n", word)
	return nil
}

func saveVocab(vocabPath, word string, entry VocabEntry) error {
	// Implement the logic to save the vocab to the specified path
	// This is a placeholder function; actual implementation will depend on your storage format (JSON, database, etc.)
	fmt.Printf("Saving vocab word %s to %s: %+v\n", word, vocabPath, entry)
	return nil
}
