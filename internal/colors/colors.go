package colors

import "fmt"

// ANSI color codes
const (
	reset = "\033[0m"
	bold  = "\033[1m"
	dim   = "\033[2m"

	// Text colors
	red     = "\033[31m"
	green   = "\033[32m"
	yellow  = "\033[33m"
	blue    = "\033[34m"
	magenta = "\033[35m"
	cyan    = "\033[36m"
	white   = "\033[37m"

	// Bright colors
	brightRed     = "\033[91m"
	brightGreen   = "\033[92m"
	brightYellow  = "\033[93m"
	brightBlue    = "\033[94m"
	brightMagenta = "\033[95m"
	brightCyan    = "\033[96m"
	brightWhite   = "\033[97m"
)

// Color wrapper functions
func Red(text string) string {
	return red + text + reset
}

func Green(text string) string {
	return green + text + reset
}

func Yellow(text string) string {
	return yellow + text + reset
}

func Blue(text string) string {
	return blue + text + reset
}

func Magenta(text string) string {
	return magenta + text + reset
}

func Cyan(text string) string {
	return cyan + text + reset
}

func White(text string) string {
	return white + text + reset
}

func BrightRed(text string) string {
	return brightRed + text + reset
}

func BrightGreen(text string) string {
	return brightGreen + text + reset
}

func BrightYellow(text string) string {
	return brightYellow + text + reset
}

func BrightBlue(text string) string {
	return brightBlue + text + reset
}

func BrightMagenta(text string) string {
	return brightMagenta + text + reset
}

func BrightCyan(text string) string {
	return brightCyan + text + reset
}

func BrightWhite(text string) string {
	return brightWhite + text + reset
}

func Bold(text string) string {
	return bold + text + reset
}

func Dim(text string) string {
	return dim + text + reset
}

// Themed functions for the vocabulary app
func Success(text string) string {
	return fmt.Sprintf("%s✅ %s%s", brightGreen, text, reset)
}

func Error(text string) string {
	return fmt.Sprintf("%s❌ %s%s", brightRed, text, reset)
}

func Warning(text string) string {
	return fmt.Sprintf("%s⚠️  %s%s", brightYellow, text, reset)
}

func Info(text string) string {
	return fmt.Sprintf("%s💡 %s%s", brightBlue, text, reset)
}

func Header(text string) string {
	return fmt.Sprintf("%s%s📚 %s%s", bold, brightCyan, text, reset)
}

func WordDisplay(text string) string {
	return fmt.Sprintf("%s%s🔷 %s%s", bold, brightMagenta, text, reset)
}

func Definition(text string) string {
	return fmt.Sprintf("%s📖 %s%s", brightBlue, text, reset)
}

func Example(text string) string {
	return fmt.Sprintf("%s💬 %s%s", brightGreen, text, reset)
}

func Stats(text string) string {
	return fmt.Sprintf("%s📊 %s%s", brightCyan, text, reset)
}

func Celebration(text string) string {
	return fmt.Sprintf("%s🎉 %s%s", brightYellow, text, reset)
}

func Prompt(text string) string {
	return fmt.Sprintf("%s%s%s", cyan, text, reset)
}

func Separator() string {
	return fmt.Sprintf("%s%s────────────────────────────────────────%s", dim, blue, reset)
}
