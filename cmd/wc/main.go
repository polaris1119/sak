package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

var (
	cFlag bool
	lFlag bool
	mFlag bool
	wFlag bool
)

func init() {
	flag.BoolVar(&cFlag, "c", false, `The number of bytes in each input file is written to the standard output.`)
	flag.BoolVar(&lFlag, "l", false, `The number of lines in each input file is written to the standard output.`)
	flag.BoolVar(&mFlag, "m", false, `The number of characters in each input file is written to the standard output.`)
	flag.BoolVar(&wFlag, "w", false, `The number of words in each input file is written to the standard output.`)
}

func main() {
	flag.Parse()

	split := bufio.ScanWords
	switch {
	case cFlag:
		split = bufio.ScanBytes
	case lFlag:
		split = bufio.ScanLines
	case mFlag:
		split = bufio.ScanRunes
	case wFlag:
		split = bufio.ScanWords
	}

	fmt.Println(count(os.Stdin, split))
}

// count 从 io.Reader 中读取数据，统计相关数据
func count(r io.Reader, split bufio.SplitFunc) int {
	scanner := bufio.NewScanner(r)

	scanner.Split(split)

	wc := 0

	for scanner.Scan() {
		wc++
	}

	return wc
}
