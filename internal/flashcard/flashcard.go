package flashcard

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/neilsmahajan/vocabulary_study_go_cli/internal/colors"
	"github.com/neilsmahajan/vocabulary_study_go_cli/internal/storage"
)

type FlashcardSession struct {
	Vocab    map[string]storage.VocabEntry
	Progress map[string]string
	Limit    int
	Review   string // "unknown", "unseen", "all"
}

func NewSession(vocab map[string]storage.VocabEntry, progress map[string]string, limit int, review string) *FlashcardSession {
	return &FlashcardSession{
		Vocab:    vocab,
		Progress: progress,
		Limit:    limit,
		Review:   review,
	}
}

func (s *FlashcardSession) Run() error {
	// Prepare unknown or unseen words
	words := []string{}
	for word := range s.Vocab {
		status := s.Progress[word]
		switch s.Review {
		case "unknown":
			if status == "unknown" {
				words = append(words, word)
			}
		case "unseen":
			if status == "" {
				words = append(words, word)
			}
		case "all":
			if status != "known" {
				words = append(words, word)
			}
		default:
			if status != "known" {
				words = append(words, word)
			}
		}
		if s.Limit > 0 && len(words) >= s.Limit {
			break
		}
	}

	// Shuffle words
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	rng.Shuffle(len(words), func(i, j int) {
		words[i], words[j] = words[j], words[i]
	})

	if len(words) == 0 {
		fmt.Println(colors.Celebration("You've marked all the words as known. Great job!"))
		return nil
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Println(colors.Header("Vocabulary Study Session"))
	fmt.Println(colors.Separator())

	for i, word := range words {
		entry := s.Vocab[word]

		// Progress indicator
		fmt.Printf("\n%s\n", colors.Dim(fmt.Sprintf("Card %d of %d", i+1, len(words))))

		// Front of card
		fmt.Printf("\n%s\n", colors.WordDisplay(word))
		fmt.Printf("  %s %s\n\n", colors.Dim("Part of Speech:"), colors.Yellow(entry.PartOfSpeech))
		fmt.Print(colors.Prompt("Press [q]uit to exit or [Enter] to flip the card... "))
		flipInput, _ := reader.ReadString('\n')
		flipInput = strings.TrimSpace(strings.ToLower(flipInput))
		if flipInput == "q" {
			fmt.Println(colors.Info("Exiting session. Your progress has been saved."))
			return nil
		}

		// Back of the card
		fmt.Println(colors.Separator())
		fmt.Printf("  %s\n", colors.Definition(entry.Definition))
		fmt.Printf("  %s\n\n", colors.Example(entry.ExampleSentence))
		fmt.Print(colors.Prompt("Did you know this word? [y]es / [n]o / [q]uit: "))

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(strings.ToLower(input))

		switch input {
		case "y":
			s.Progress[word] = "known"
			fmt.Println(colors.Success("Marked as known!"))
		case "n":
			s.Progress[word] = "unknown"
			fmt.Println(colors.Warning("Marked for review."))
		case "q":
			fmt.Println(colors.Info("Exiting session. Your progress has been saved."))
			return nil
		default:
			fmt.Println(colors.Error("Invalid input. Skipping word."))
		}

		if i < len(words)-1 {
			fmt.Printf("\n%s\n", colors.Separator())
		}
	}
	fmt.Printf("\n%s\n", colors.Success("End of session! Progress saved."))
	return nil
}
