package main

import (
	"fmt"
	"log"
	//"github.com/neilsmahajan/vocabulary_study_go_cli/internal/cli"
	//"os"

	"github.com/neilsmahajan/vocabulary_study_go_cli/internal/storage"
)

func main() {
	//err := cli.Run()

	vocabPath := "vocab.json"
	vocabMap, err := storage.LoadVocab(vocabPath)
	if err != nil {
		log.Fatalf("Could not load vocab: %v", err)
	}
	fmt.Println("loaded ", len(vocabMap), "words")
	for key, value := range vocabMap {
		fmt.Println("Word: ", key, " POS: ", value.PartOfSpeech, " Definition: ", value.Definition, " Example: ", value.ExampleSentence)
	}
}
