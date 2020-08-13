### Steps to execute:

Change to the cmd directory

go build

Run the executable generated, this would execute the main function which accepts a file path.

Note: The code works for go 1.8 and above 

### Tests:

I have written unit test cases as well just to be sure that the code is working properly.

**Data Structure used to hold word data:**
1. Binary Search Tree: Data is stored alphabetically. This gives O(logN) insertion and lookup time, but O(NlogN) time
to fetch the top 20 repeated elements.
2. WordCountStoreEntry: This struct is used to hold key value pairs, ie. word and its corresponding count, this struct is stored in the BST.
   
**Strings vs Byte Array for Comparison Operation**
Strings are faster than byte array for comparisons, that is why, strings are used to store a word in WordCountStoreEntry. https://medium.com/@felipedutratine/in-golang-should-i-work-with-bytes-or-strings-8bd1f5a7fd48


**Main Components**
1. File Injestor: Loads the file into a text string.
2. TextToWordsConverter: Transforms that string into a WordCountStore.
3. WordCountStore: The struct which wraps the BST which holds the complete data set.

### Improvements
A combination of BST and heap can be used instead of plain BST, this would make fetching the top 20 elements possible in logarithmic time complexity. 
