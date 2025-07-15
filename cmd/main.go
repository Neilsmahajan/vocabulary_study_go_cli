package main

import (
	"fmt"
	"os"

	"github.com/neilsmahajan/vocabulary_study_go_cli/internal/cli"
)

func main() {
	if err := cli.Run(); err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
}
