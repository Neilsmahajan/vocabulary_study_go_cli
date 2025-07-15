package cli

import (
	"fmt"

	"github.com/neilsmahajan/vocabulary_study_go_cli/internal/flashcard"
	"github.com/neilsmahajan/vocabulary_study_go_cli/internal/storage"
)

func Run() error {
	const vocabPath = "vocab.json"
	const progressPath = "progress.json"

	vocab, err := storage.LoadVocab(vocabPath)
	if err != nil {
		return fmt.Errorf("error loading vocab: %w", err)
	}

	progress, err := storage.LoadProgress(progressPath)
	if err != nil {
		return fmt.Errorf("error loading progress: %w", err)
	}

	// Start flashcard session
	session := flashcard.NewSession(vocab, progress)
	if err := session.Run(); err != nil {
		return fmt.Errorf("error running flashcard session: %w", err)
	}

	// Save progress
	if err := storage.SaveProgress(progressPath, progress); err != nil {
		return fmt.Errorf("error saving progress: %w", err)
	}
	return nil
}
