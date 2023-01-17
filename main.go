package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	// define flags
	flag.Bool("l", false, "count lines")
	flag.Bool("b", false, "count bytes")

	flag.Parse()
	usedFlags := make(map[string]bool)
	flag.Visit(func(f *flag.Flag) {
		usedFlags[f.Name] = true
	})

	// Calling the count function to count the number of words
	// received from the Standard Input and printing it out
	fmt.Println(count(os.Stdin, usedFlags["l"], usedFlags["b"]))
}

func count(r io.Reader, countLines, countBytes bool) int {
	// A scanner is used to read text from a Reader (such as files)
	scanner := bufio.NewScanner(r)

	// If the count lines flag is not set, we want to count words os we define
	// the scanner split type to words (default is split by lines)
	if !countLines {
		// Define the scanner split type to words(default is split by lines)
		scanner.Split(bufio.ScanWords)
	}

	// If the countBytes flag is set, we use the Bytes() method from teh bufio scanner struct
	// to count bytes when the scanner.Scan() is looped over
	if countBytes {
		// split with bufio.ScanBytes function to count how many bytes there are
		scanner.Split(bufio.ScanBytes)
	}

	// Defining a counter
	wc := 0

	// For every word or line scanned, increment the counter
	for scanner.Scan() {
		wc++
	}

	// Return the total
	return wc
}
