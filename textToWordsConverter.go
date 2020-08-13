package word

import "strings"

// TextToWordsConverter - Interface which processeses the entire string and
// gives back the words contained in it wrapped in a WordCountStore
type TextToWordsConverter interface {
	ProcessAndConvertToWordCountStore(text string) WordCountStore
}

type textToWordsConverter struct{}

// NewTextToWordsConverter Constructor for TextToWordsConverter
func NewTextToWordsConverter() TextToWordsConverter {
	return &textToWordsConverter{}
}

func splitFuncModern(r rune) bool {
	if r >= 'a' && r <= 'z' {
		return false
	} else if r >= 'A' && r <= 'Z' {
		return false
	} else {
		return true
	}
}
func (conveter *textToWordsConverter) ProcessAndConvertToWordCountStore(text string) WordCountStore {
	// A key advantage with Fields: we treat many consecutive delimiter characters as one.
	// So several spaces is the same as a single space. This is a key advantage of Fields over Split().
	words := strings.FieldsFunc(text, splitFuncModern)
	wordCountStore := NewBSTBackedWordCountStore()
	for _, word := range words {
		wordCountStore.AddOrIncrementCount(strings.ToLower(word))
	}
	return wordCountStore
}
