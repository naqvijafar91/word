package word

import "sort"

// WordCountStoreEntry represents each word and its count
type WordCountStoreEntry struct {
	Word  string
	Count int
}

func createWordCountStoreEntry(word string) *WordCountStoreEntry {
	return &WordCountStoreEntry{word, 1}
}

type WordCountStore interface {
	AddOrIncrementCount(word string)
	GetTopFrequentWords() []*WordCountStoreEntry
	GetCountForWord(word string) int
}

type treeNode struct {
	key   *WordCountStoreEntry
	left  *treeNode
	right *treeNode
}

// BSTBackedWordCountStore uses a Binary Search Tree to store a word count pair, the tree is sorted alphabetically.
// With the help of this tree, we have O(logN) insertion and lookup time. Fetching top 20 elements takes  O(NlogN) time
// but since we are only doing it once, that is acceptable. A treap can be used to improve that as well.
// ---------------------------------------------------------------------
// We can use a sorted slice based word count store, but that would give O(n) insertion time: O(logN) for finding the
// spot to insert and O(n) for shifting the remaining elements, so total O(n)
// ---------------------------------------------------------------------
// We can alse use a linked list, this would require no shifting and anything can be inserted in between, but it requires
// O(N) time to find the insertion spot.
// ---------------------------------------------------------------------
// NOTE: Using a map is the direct approach for this kind of scenario, but since we cannot use it for this test, I have
// implemented via BST which is the next best thing. I could have implemented my own map from scratch, but I think
// that is not something which was asked out of this assignment.
type BSTBackedWordCountStore struct {
	root *treeNode
}

func NewBSTBackedWordCountStore() WordCountStore {
	return &BSTBackedWordCountStore{}
}

func (store *BSTBackedWordCountStore) AddOrIncrementCount(word string) {
	searchedEntry := store.binarySearch(store.root, word)
	if searchedEntry != nil {
		searchedEntry.Count++
	} else {
		// Insert such that the linked list is sorted
		store.root = store.insertIntoBSTRecursively(store.root, createWordCountStoreEntry(word))
	}
}

func (store *BSTBackedWordCountStore) binarySearch(root *treeNode, word string) *WordCountStoreEntry {
	if root == nil {
		return nil
	}
	if root.key.Word == word {
		return root.key
	}
	// val is greater than root's key
	if root.key.Word > word {
		return store.binarySearch(root.left, word)
	}
	// val is less than root's key
	return store.binarySearch(root.right, word)
}

func (store *BSTBackedWordCountStore) insertIntoBSTRecursively(root *treeNode, entry *WordCountStoreEntry) *treeNode {
	if root == nil {
		root = &treeNode{key: entry}
		return root
	}
	if entry.Word < root.key.Word {
		root.left = store.insertIntoBSTRecursively(root.left, entry)
	} else {
		root.right = store.insertIntoBSTRecursively(root.right, entry)
	}
	return root
}

// Note: The binary tree is sorted alphabetically, so we need to extract out everything and sort by word count
func (store *BSTBackedWordCountStore) GetTopFrequentWords() []*WordCountStoreEntry {
	sortedSlice := make([]*WordCountStoreEntry, 0)
	sortedSlice = store.extractToSlice(store.root, sortedSlice)
	sort.Slice(sortedSlice, func(i, j int) bool {
		return sortedSlice[i].Count > sortedSlice[j].Count
	})
	// We only need top 20
	if(len(sortedSlice)>20) {
		sortedSlice = sortedSlice[0:20]
	}
	return sortedSlice
}

func (store *BSTBackedWordCountStore) extractToSlice(root *treeNode, slice []*WordCountStoreEntry) []*WordCountStoreEntry {
	if root == nil {
		return slice
	}
	slice = append(slice, root.key)
	slice = store.extractToSlice(root.left, slice)
	slice = store.extractToSlice(root.right, slice)
	return slice
}

func (store *BSTBackedWordCountStore) GetCountForWord(word string) int {
	entry := store.binarySearch(store.root, word)
	if entry == nil {
		return -1
	}
	return entry.Count
}
