package cli

import (
	"fmt"
	"os"
	"strings"

	"github.com/neilsmahajan/vocabulary_study_go_cli/internal/flashcard"
	"github.com/neilsmahajan/vocabulary_study_go_cli/internal/storage"
)

func Run() error {
	if len(os.Args) > 1 && strings.ToLower(os.Args[1]) == "stats" {
		return showStats()
	}
	if len(os.Args) > 1 && strings.ToLower(os.Args[1]) == "reset" {
		return resetProgress()
	}
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

func showStats() error {
	const vocabPath = "vocab.json"
	const progressPath = "progress.json"

	vocab, err := storage.LoadVocab(vocabPath)
	if err != nil {
		return err
	}
	progress, err := storage.LoadProgress(progressPath)
	if err != nil {
		return err
	}

	total := len(vocab)
	known, unknown := 0, 0

	for word := range vocab {
		status := progress[word]
		switch status {
		case "known":
			known++
		case "unknown":
			unknown++
		}
	}
	unseen := total - known - unknown

	fmt.Printf("\n📊 Vocabulary Stats:\n")
	fmt.Printf("Total words: %d\n", total)
	fmt.Printf("Known: %d\n", known)
	fmt.Printf("Unknown: %d\n", unknown)
	fmt.Printf("Unseen: %d\n", unseen)
	return nil
}

func resetProgress() error {
	const progressPath = "progress.json"

	fmt.Print("⚠️ Are you sure you want to reset your progress? [y/N]: ")
	var response string
	fmt.Scanln(&response)

	if strings.ToLower(response) != "y" {
		fmt.Println("❎ Reset cancelled.")
		return nil
	}
	err := storage.SaveProgress(progressPath, map[string]string{})
	if err != nil {
		return fmt.Errorf("failed to reset progress: %w", err)
	}
	fmt.Println("✅ Progress has been reset.")
	return nil
}
