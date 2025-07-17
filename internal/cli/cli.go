package cli

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/neilsmahajan/vocabulary_study_go_cli/internal/flashcard"
	"github.com/neilsmahajan/vocabulary_study_go_cli/internal/storage"
)

func Run() error {
	// Handle subcommands
	if len(os.Args) > 1 && !strings.HasPrefix(os.Args[1], "-") {
		arg := os.Args[1]
		switch strings.ToLower(arg) {
		case "stats":
			return showStats()
		case "reset":
			return resetProgress()
		case "help":
			printUsage()
			return nil
		}
	}

	// Flag parsing
	limit := flag.Int("limit", 0, "Maximum number of flashcards in the session")
	review := flag.String("review", "all", "Choose which words to review: all, unknown, unseen")
	help := flag.Bool("help", false, "Show help message")
	flag.Usage = printUsage
	flag.Parse()

	if *help {
		printUsage()
		return nil
	}

	// Validate review mode
	if *review != "all" && *review != "unknown" && *review != "unseen" {
		fmt.Println("❌ Invalid value for --review. Use: all, unknown, or unseen")
		printUsage()
		return nil
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
	session := flashcard.NewSession(vocab, progress, *limit, *review)
	if err := session.Run(); err != nil {
		return fmt.Errorf("error running flashcard session: %w", err)
	}

	return storage.SaveProgress(progressPath, session.Progress)
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
	if total == 0 {
		fmt.Println("No words in vocabulary.")
	}
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

func printUsage() {
	fmt.Println(`
📚 Vocabulary Study CLI

Usage:
  vocab [flags]
  vocab stats         Show study statistics
  vocab reset         Reset all progress
  vocab help          Show this help message

Flags:
  --limit N           Limit number of flashcards shown in one session
  --review MODE       Filter words to review: all, unknown, unseen (default: all)
  --help              Show this help message

Examples:
  vocab --limit 20
  vocab --review=unknown
  vocab stats`)
}
