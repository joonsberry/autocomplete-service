package main

import (
	"log"
	"strings"
	"sync"
)

// UTILITY FUNCTIONS

// check error and panic if not nil
func check(e error) {
	if e != nil {
		panic(e)
	}
}

// adjustBuffer ensures that the next buffer to be searched does not cut off words
// returns a complete buffer and what the next seek value should be
func adjustBuffer(buf []byte, seek int64, fsize int64) (chunk []byte, newSeek int64) {

	passage := string(buf) // convert []byte to string

	// check whether to cutoff at a space or a newline
	cutoff := strings.LastIndexByte(passage, ' ')
	cutoffNL := strings.LastIndexByte(passage, '\n')
	if cutoffNL > cutoff {
		cutoff = cutoffNL
	}

	// if cutoff == -1, the chunking is too small, need to reduce granularity
	if cutoff == -1 {
		log.Println(passage)
		log.Fatal("Reduce concurrency threshold")
	}

	cutoff++ // increment to get correct high bound

	// return the complete buffer and the next seek value
	return []byte(passage[0:cutoff]), seek + int64(cutoff)
}

// searchChunk takes one chunk of the input text, maps terms matching the query, and adds this map to a channel
func searchChunk(chunk []byte, query string, wg *sync.WaitGroup, results chan<- map[string]int) {

	defer wg.Done() // decrement wg when routine returns
	res := make(map[string]int)

	passage := string(chunk)             // convert []byte to string
	terms := strings.Split(passage, " ") // split passage on spaces

	// loop through all terms in passage
	for _, term := range terms {

		// trim non-alpha chars
		t := strings.TrimFunc(term, func(c rune) bool {
			return !(c >= 65 && c <= 90) && !(c >= 97 && c <= 122)
		})

		t = strings.ToLower(t) // convert to lowercase

		// either add map entry or increment current entry
		if strings.Contains(t, query) {
			if _, ok := res[t]; ok {
				res[t]++
			} else {
				res[t] = 1
			}
		}
	}

	results <- res // send map to channel
}

// reduceChuncks maps all frequency maps into one map and returns a sorted PairList
func reduceChunks(results <-chan map[string]int) PairList {

	wordfreqs := make(map[string]int)

	// loop through maps on the results channel
	for res := range results {

		// loop through terms in each map to add or increment the master map
		for term, freq := range res {
			if _, ok := wordfreqs[term]; ok {
				wordfreqs[term] += freq
			} else {
				wordfreqs[term] = freq
			}
		}
	}

	return rankByWordCount(wordfreqs) // sort word frequencies and return
}
