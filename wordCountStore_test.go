package word

import "testing"

func TestAddOrIncrementCountPart1(t *testing.T) {
	store := NewBSTBackedWordCountStore()
	store.AddOrIncrementCount("jafar")
	if store.GetCountForWord("jafar") != 1 {
		t.Errorf("Count should be 1")
	}
	store.AddOrIncrementCount("jafar")
	if store.GetCountForWord("jafar") != 2 {
		t.Errorf("Count should be 2")
	}
}

func TestAddOrIncrementCountPart2(t *testing.T) {
	store := NewBSTBackedWordCountStore()
	for i := 0; i < 14; i++ {
		store.AddOrIncrementCount("jafar")
	}
	if store.GetCountForWord("jafar") != 14 {
		t.Errorf("Count should be 14")
	}
}

func TestGetTopFrequentWords(t *testing.T) {
	store := NewBSTBackedWordCountStore()
	addWordWithCount(store, "jafar", 101)
	addWordWithCount(store, "xyz", 1011)
	addWordWithCount(store, "juy", 4)
	resultListSorted := store.GetTopFrequentWords()
	if len(resultListSorted) != 3 {
		t.Error("Length incorrect")
	}
	if resultListSorted[0].Word != "xyz" && resultListSorted[0].Count != 1011 {
		t.Error("Word or count mismatch")
	}
	if resultListSorted[1].Word != "jafar" && resultListSorted[0].Count != 101 {
		t.Error("Word or count mismatch")
	}
	if resultListSorted[2].Word != "juy" && resultListSorted[0].Count != 4 {
		t.Error("Word or count mismatch")
	}
}

func addWordWithCount(store WordCountStore, word string, count int) {
	for i := 0; i < count; i++ {
		store.AddOrIncrementCount(word)
	}
}
