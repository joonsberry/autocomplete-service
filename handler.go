package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
)

// HANDLERS

// autocompleteHandler handles requests to /autocomplete
func autocompleteHandler(w http.ResponseWriter, r *http.Request) {

	log.Print("HANDLER START --------------------\n\n")

	start := time.Now() // used to time handler

	// verify the request uses GET method
	if r.Method != "GET" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad Request: resource only supports GET"))
		return
	}

	query, ok := r.URL.Query()["term"] // get term query param

	// check that term query param exists and that there is only one
	if !ok || len(query) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad Request: missing term parameter"))
		return
	} else if len(query) > 1 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad Request: provide only one term parameter"))
		return
	}

	term := query[0] // get string value from query array

	// open input file to be searched
	f, err := os.Open("input.txt")
	check(err)

	// get file info
	finfo, err := f.Stat()
	check(err)

	// use file size to calculate standard buffer size
	fsize := finfo.Size()
	bufsize := fsize / 30

	log.Print("File Size: ", fsize, "\n\n")

	// initialize chunk []byte and seek counter
	var chunk []byte
	seek := int64(0)

	// initialize waitgroup for goroutines and the results channel (at most fsize sends)
	var wg sync.WaitGroup
	results := make(chan map[string]int, fsize)

	// chunk though buffer until seek has reached EOF
	for seek < fsize {

		wg.Add(1) // increment wg

		// ensure buffer does not attempt to read past EOF
		if seek+bufsize > fsize {
			bufsize = fsize - seek
		}

		// read next chunk into buffer
		buf := make([]byte, bufsize)
		_, err := f.ReadAt(buf, seek)
		check(err)

		chunk, seek = adjustBuffer(buf, seek, fsize) // get complete chunk

		go searchChunk(chunk, term, &wg, results) // launch goroutine to search this chunk and continue

	}

	// wait until all chunks have been analyzed and close channel
	wg.Wait()
	close(results)

	final := reduceChunks(results) // reduce chunks into sorted PairList

	// log 25 most frequent terms and build response string
	resp := "25 Most Frequent Words Including the Term: " + term + "\n"
	i := 0
	for i < 25 {
		log.Println(final[i])
		resp += final[i].Key + ": " + strconv.Itoa(final[i].Value) + "\n"
		i++
	}

	// write success response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(resp))

	// log elapsed time for handler and return
	elapsed := time.Since(start)
	log.Print("HANDLER DONE IN ", elapsed, "----------\n\n\n")
	return
}
