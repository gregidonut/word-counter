package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	// Defining a boolean flag -l to count lines instead of words
	lines := flag.Bool("l", false, "count lines")

	// Defining a boolean flag -b to count bytes
	bytes := flag.Bool("b", false, "count lines")

	// Parsing the flags provided by the user
	flag.Parse()

	// Calling the count function to count the number of words
	// received from the Standard Input and printing it out
	fmt.Println(count(os.Stdin, *lines, *bytes))
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
		scanner.Bytes()
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
