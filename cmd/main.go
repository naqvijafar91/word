package main

import (
	"fmt"
	"log"
	"os"
	"github.com/naqvijafar91/word"
)

func main() {
	filePath := "./sample/mobydick.txt"
	if len(os.Args[1:]) == 1 {
		filePath = os.Args[1:][0]
	} else {
		log.Printf("No file path specified, using default file at %s for this execution", filePath)
	}
	textInByteArray, err := word.NewFileInjestor().Injest(filePath)
	if err != nil {
		log.Fatal("Unable to read file", err)
	}
	textToWordsConverter := word.NewTextToWordsConverter()
	wordCountStore := textToWordsConverter.ProcessAndConvertToWordCountStore(string(textInByteArray))
	wordsWithCount := wordCountStore.GetTopFrequentWords()
	for _, entry := range wordsWithCount {
		fmt.Printf("%s : %d \n", entry.Word, entry.Count)
	}
}
