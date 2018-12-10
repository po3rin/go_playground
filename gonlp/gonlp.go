package gonlp // import "github.com/po3rin/gonlp"

import (
	"strings"
)

// Corpus type is include id num only.
type Corpus []int

// WordToID for changing word to id.
type WordToID map[string]int

// IDToWord for changing id to word.
type IDToWord map[int]string

// NewProcess create corpus, wordToID, idToWprd.
func NewProcess(text string) (Corpus, WordToID, IDToWord) {
	text = strings.ToLower(text)
	text = strings.Replace(text, ".", " .", -1)
	words := strings.Split(text, " ")

	var corpus Corpus
	wordToID := make(WordToID)
	idToWord := make(IDToWord)

	for _, word := range words {
		if !containsString(wordToID, word) {
			newID := len(wordToID)
			wordToID[word] = newID
			idToWord[newID] = word
		}
	}

	for _, word := range words {
		corpus = append(corpus, wordToID[word])
	}

	return corpus, wordToID, idToWord
}
