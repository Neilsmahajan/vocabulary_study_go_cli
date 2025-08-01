package flashcard

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

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
		fmt.Println("🎉 You've marked all the words as known. Great job!")
		return nil
	}

	reader := bufio.NewReader(os.Stdin)
	for _, word := range words {
		entry := s.Vocab[word]

		// Front of card
		fmt.Printf("\n🔷 Word: %s\nPart of Speech: %s\n", word, entry.PartOfSpeech)
		fmt.Print("Press [q]uit to exit or [Enter] to flip the card...")
		flipInput, _ := reader.ReadString('\n')
		flipInput = strings.TrimSpace(strings.ToLower(flipInput))
		if flipInput == "q" {
			fmt.Println("👋 Exiting session. Your progress has been saved.")
			return nil
		}
		// Back of the card
		fmt.Printf("\n📖 Definition: %s\n", entry.Definition)
		fmt.Printf("💬 Example: %s\n", entry.ExampleSentence)
		fmt.Print("Did you know this word? [y]es / [n]o / [q]uit: ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(strings.ToLower(input))

		switch input {
		case "y":
			s.Progress[word] = "known"
		case "n":
			s.Progress[word] = "unknown"
		case "q":
			fmt.Println("👋 Exiting session. Your progress has been saved.")
			return nil
		default:
			fmt.Println("‼️ Invalid input. Skipping word.")
		}
	}
	fmt.Println("\n✅ End of session! Progress saved.")
	return nil
}
