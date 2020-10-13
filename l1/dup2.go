package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// dup1 prints duplicated lines to stdout preceded by line count
// reads from stdinput or from a list of files
func main() {
	counts := make(map[string]int)
	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, counts, "os.Stdin")
	} else {
		for _, fName := range files {
			f, err := os.Open(fName)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}

			countLines(f, counts, fName)
			f.Close()
		}
	}

	printLines(counts)
}

func printLines(counts map[string]int) {
	for line, n := range counts {
		if n > 1 {
			fileNameAdLine := strings.Split(line, "#")
			fileName, l := fileNameAdLine[0], fileNameAdLine[1]
			fmt.Printf("%d\t%s\t%s\n", n, fileName, l)
		}
	}
}

func countLines(f *os.File, counts map[string]int, fileName string) {
	input := bufio.NewScanner(f)

	for input.Scan() {
		counts[fileName+"#"+input.Text()]++
	}
}
