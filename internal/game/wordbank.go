package game

import (
	"math/rand"
)

type WordBank struct {
	Words map[int][]string
}

func NewWordBank() *WordBank {
	return &WordBank{
		Words: make(map[int][]string),
	}
}

func (wb *WordBank) AddWord(word string, difficulty int) {
	wb.Words[difficulty] = append(wb.Words[difficulty], word)
}

func (wb *WordBank) GetWord(difficulty int) string {
	words := wb.Words[difficulty]
	if len(words) == 0 {
		return ""
	}
	return words[rand.Intn(len(words))]
}
