package main

import (
	"sort"
)

// the following sorting interface was taken from:
// https://stackoverflow.com/questions/18695346/how-to-sort-a-mapstringint-by-its-values
// comments are my own

// rankByWordCount takes a map of word frequencies and returns a sorted PairList
func rankByWordCount(wordFrequencies map[string]int) PairList {

	pl := make(PairList, len(wordFrequencies))
	i := 0
	for k, v := range wordFrequencies {
		pl[i] = Pair{k, v}
		i++
	}
	sort.Sort(sort.Reverse(pl))
	return pl
}

// Pair stores a key/value pair
type Pair struct {
	Key   string
	Value int
}

type PairList []Pair

// the following functions allow PaitList to implement sort interface
func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
