package main

import (
	"fmt"
	"sort"
	"strings"
)

type Word struct {
	Word      string
	Frequency int
}

type ByFrequency []Word

func (p ByFrequency) Len() int {
	return len(p)
}

func (p ByFrequency) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p ByFrequency) Less(i, j int) bool {
	return p[i].Frequency > p[j].Frequency
}

func SortByFrequency(words []Word) []Word {
	sort.Sort(ByFrequency(words))
	return words
}

func CountWords(text string) map[string]int {
	words := strings.Fields(text)
	wordCount := make(map[string]int)

	for _, word := range words {
		word = strings.Trim(word, ".,!?:;\"'")
		if _, ok := wordCount[word]; ok {
			wordCount[word]++
		} else {
			wordCount[word] = 1
		}
	}
	return wordCount
}

func Top5Words(wordmap map[string]int) []Word {
	var words []Word
	for text, count := range wordmap {
		words = append(words, Word{Word: text, Frequency: count})
	}

	SortByFrequency(words)

	if len(words) > 5 {
		return words[:5]
	}
	return words
}

func main() {
	fmt.Println("Word Frequency Test")

	text := `Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.`

	results := CountWords(text)
	MostCommon := Top5Words(results)

	fmt.Println(MostCommon)
}
