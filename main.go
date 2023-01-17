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
	list := flag.Bool("l", false, "count lines")
	bytes := flag.Bool("b", false, "count bytes")
	file := flag.Bool("file", false, "Specify file name to count words from")

	flag.Parse()

	files := make([]string, 0)
	if *file {
		files = append(files, flag.Args()...)
	}

	if err := run(*list, *bytes, files); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(countLines, countBytes bool, files []string) error {
	// injecting dependency for easier testing
	return actualRun(os.Stdout, countLines, countBytes, files)
}

func actualRun(w io.Writer, countLines, countBytes bool, files []string) error {
	// should check if there are any arguments at
	// all coz program cant do anything without nothing to count
	if len(os.Args) == 1 {
		// this message should be flags.Usage()
		// but temporarily a one line string to simplify testing
		return fmt.Errorf("please provide arguments")
	}

	if len(files) > 0 {
		total := 0
		for _, v := range files {
			fileBytes, err := os.OpenFile(v, os.O_RDONLY, 0755)
			if err != nil {
				return err
			}

			byteCount := count(fileBytes, countLines, countBytes)
			total += byteCount
			fmt.Fprintf(w, "%s: %d\n", v, byteCount)
		}
		fmt.Fprintf(w, "total: %d\n", total)
		return nil
	}

	fmt.Fprintln(w, count(os.Stdin, countLines, countBytes))
	return nil
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
