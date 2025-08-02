package cli

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/neilsmahajan/vocabulary_study_go_cli/internal/colors"
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
		case "add":
			return handleAddCommand()
		case "remove":
			return handleRemoveCommand()
		default:
			fmt.Printf("%s\n", colors.Error(fmt.Sprintf("Unknown command: %s", arg)))
			fmt.Println(colors.Info("Use './vocab help' to see available commands."))
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
		fmt.Println(colors.Error("Invalid value for --review. Use: all, unknown, or unseen"))
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

	fmt.Printf("\n%s\n", colors.Stats("Vocabulary Stats"))
	fmt.Println(colors.Separator())
	fmt.Printf("%s %s\n", colors.Dim("Total words:"), colors.BrightWhite(fmt.Sprintf("%d", total)))
	fmt.Printf("%s %s\n", colors.Dim("Known:"), colors.BrightGreen(fmt.Sprintf("%d", known)))
	fmt.Printf("%s %s\n", colors.Dim("Unknown:"), colors.BrightRed(fmt.Sprintf("%d", unknown)))
	fmt.Printf("%s %s\n", colors.Dim("Unseen:"), colors.BrightYellow(fmt.Sprintf("%d", unseen)))

	if total == 0 {
		fmt.Printf("\n%s\n", colors.Warning("No words in vocabulary."))
	} else {
		percentage := float64(known) / float64(total) * 100
		progressBar := generateProgressBar(percentage, 20)
		fmt.Printf("\n%s %.1f%% %s\n", colors.Info("Progress:"), percentage, progressBar)
	}
	fmt.Println()
	return nil
}

func generateProgressBar(percentage float64, width int) string {
	filled := int(percentage / 100 * float64(width))
	bar := ""
	for i := 0; i < width; i++ {
		if i < filled {
			bar += colors.BrightGreen("█")
		} else {
			bar += colors.Dim("░")
		}
	}
	return fmt.Sprintf("[%s]", bar)
}

func resetProgress() error {
	const progressPath = "progress.json"

	fmt.Print(colors.Warning("Are you sure you want to reset your progress? [y/N]: "))
	var response string
	if _, err := fmt.Scanln(&response); err != nil {
		// Handle empty input or scan errors gracefully
		response = "n"
	}

	if strings.ToLower(response) != "y" {
		fmt.Println(colors.Info("Reset canceled."))
		return nil
	}
	err := storage.SaveProgress(progressPath, map[string]string{})
	if err != nil {
		return fmt.Errorf("failed to reset progress: %w", err)
	}
	fmt.Println(colors.Success("Progress has been reset."))
	return nil
}

func handleAddCommand() error {
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)

	wordFlag := addCmd.String("word", "", "Word to add")
	posFlag := addCmd.String("pos", "", "Part of speech (noun, verb, adjective, etc.)")
	definitionFlag := addCmd.String("definition", "", "Definition of the word")
	exampleFlag := addCmd.String("example", "", "Example sentence using the word")

	if err := addCmd.Parse(os.Args[2:]); err != nil {
		return err
	}

	if *wordFlag == "" || *posFlag == "" || *definitionFlag == "" || *exampleFlag == "" {
		fmt.Println(colors.Error("All flags (--word, --pos, --definition, --example) are required."))
		fmt.Printf("%s %s\n", colors.Dim("Example:"), colors.Cyan("./vocab add --word=précis --pos=noun --definition=\"a summary or abstract of a text or speech\" --example=\"You can read a brief precis of what he found by clicking here.\""))
		fmt.Println(colors.Info("Use './vocab help' to see available commands."))
		return nil
	}

	word := *wordFlag
	pos := *posFlag
	definition := *definitionFlag
	example := *exampleFlag

	fmt.Printf("\n%s\n", colors.Header("Adding New Word"))
	fmt.Println(colors.Separator())
	fmt.Printf("%s %s\n", colors.Dim("Word:"), colors.BrightMagenta(word))
	fmt.Printf("%s %s\n", colors.Dim("Part of speech:"), colors.Yellow(pos))
	fmt.Printf("%s %s\n", colors.Dim("Definition:"), colors.BrightBlue(definition))
	fmt.Printf("%s %s\n", colors.Dim("Example:"), colors.BrightGreen(example))
	fmt.Println()

	if err := storage.AddWord("vocab.json", word, pos, definition, example); err != nil {
		return fmt.Errorf("failed to add word: %w", err)
	}

	return nil
}

func handleRemoveCommand() error {
	removeCmd := flag.NewFlagSet("remove", flag.ExitOnError)

	wordFlag := removeCmd.String("word", "", "Word to remove")

	if err := removeCmd.Parse(os.Args[2:]); err != nil {
		return err
	}

	if *wordFlag == "" {
		fmt.Println(colors.Error("The --word flag is required to remove a word."))
		fmt.Printf("%s %s\n", colors.Dim("Example:"), colors.Cyan("./vocab remove --word=précis"))
		fmt.Println(colors.Info("Use './vocab help' to see available commands."))
		return nil
	}

	word := *wordFlag

	fmt.Printf("\n%s\n", colors.Header("Removing Word"))
	fmt.Println(colors.Separator())
	fmt.Printf("%s %s\n", colors.Dim("Word to remove:"), colors.BrightRed(word))
	fmt.Println()

	if err := storage.RemoveWord("vocab.json", word); err != nil {
		return fmt.Errorf("failed to remove word: %w", err)
	}
	return nil
}

func printUsage() {
	fmt.Printf("\n%s\n", colors.Header("Vocabulary Study CLI"))
	fmt.Println(colors.Separator())

	fmt.Printf("\n%s\n", colors.Bold("Usage:"))
	fmt.Printf("  %s\n", colors.Cyan("./vocab [flags]"))
	fmt.Printf("  %s\t\t%s\n", colors.Cyan("./vocab stats"), colors.Dim("Show study statistics"))
	fmt.Printf("  %s\t\t%s\n", colors.Cyan("./vocab reset"), colors.Dim("Reset all progress"))
	fmt.Printf("  %s\t%s\n", colors.Cyan("./vocab add --word=<word> --pos=<part-of-speech> --definition=\"<definition>\" --example=\"<example>\""), colors.Dim("Add a new word"))
	fmt.Printf("  %s\t%s\n", colors.Cyan("./vocab remove --word=<word>"), colors.Dim("Remove a word from vocabulary"))
	fmt.Printf("  %s\t\t%s\n", colors.Cyan("./vocab help"), colors.Dim("Show this help message"))

	fmt.Printf("\n%s\n", colors.Bold("Flags:"))
	fmt.Printf("  %s\t\t%s\n", colors.Yellow("--limit N"), colors.Dim("Limit number of flashcards shown in one session"))
	fmt.Printf("  %s\t%s\n", colors.Yellow("--review MODE"), colors.Dim("Filter words to review: all, unknown, unseen (default: all)"))
	fmt.Printf("  %s\t\t%s\n", colors.Yellow("--help"), colors.Dim("Show this help message"))

	fmt.Printf("\n%s\n", colors.Bold("Examples:"))
	fmt.Printf("  %s\n", colors.BrightGreen("./vocab --limit 20"))
	fmt.Printf("  %s\n", colors.BrightGreen("./vocab --review=unknown"))
	fmt.Printf("  %s\n", colors.BrightGreen("./vocab reset"))
	fmt.Printf("  %s\n", colors.BrightGreen("./vocab stats"))
	fmt.Printf("  %s\n", colors.BrightGreen("./vocab add --word=précis --pos=noun --definition=\"a summary or abstract\" --example=\"Read the précis.\""))
	fmt.Printf("  %s\n", colors.BrightGreen("./vocab remove --word=précis"))
	fmt.Println()
}
