package word

import (
	"testing"
)

func TestShouldConvertParagraph(t *testing.T) {
	text := `Contrary to popular belief, Lorem Ipsum is not simply random text. It has roots in a piece of 
	classical Latin literature from 45 BC, making it over 2000 years old. Richard McClintock, a 
	Latin professor at Hampden-Sydney College in Virginia, looked up one of the more obscure Latin words, 
	consectetur, from a Lorem Ipsum passage, and going through the cites of the word in classical literature, 
	discovered the undoubtable source. Lorem Ipsum comes from sections 1.10.32 and 1.10.33 of "de Finibus 
	Bonorum et Malorum" (The Extremes of Good and Evil) by Cicero, written in 45 BC. This book is a treatise on
	 the theory of ethics, very popular during the Renaissance. The first line of Lorem Ipsum, "Lorem 
	 ipsum dolor sit amet..", comes from a line in section 1.10.32.`
	textToWordsConverter := NewTextToWordsConverter()
	store := textToWordsConverter.ProcessAndConvertToWordCountStore(text)
	if store.GetCountForWord("of") != 7 {
		t.Errorf("Count should have been %d but it is %d", 2, store.GetCountForWord("my"))
	}
	if store.GetCountForWord("the") != 8 {
		t.Errorf("Count should have been %d but it is %d", 1, store.GetCountForWord("the"))
	}
	if store.GetCountForWord("literature") != 2 {
		t.Errorf("Count should have been %d but it is %d", 1, store.GetCountForWord("warmth"))
	}
	if store.GetCountForWord("has") != 1 {
		t.Errorf("Count should have been %d but it is %d", 1, store.GetCountForWord("you"))
	}
	if store.GetCountForWord("lorem") != 5 {
		t.Errorf("Count should have been %d but it is %d", 1, store.GetCountForWord("in"))
	}
}
