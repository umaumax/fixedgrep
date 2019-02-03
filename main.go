package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	prefix string
	suffix string
	nMax   int
)

func init() {
	flag.StringVar(&prefix, "prefix", "", "fixed prefix string")
	flag.StringVar(&suffix, "suffix", "", "fixed suffix string")
	flag.IntVar(&nMax, "max", -1, "max hit number")
}

func main() {
	flag.Parse()
	scanner := bufio.NewScanner(os.Stdin)

	keywords := flag.Args()
	if nMax == 0 {
		return
	}

	cnt := 0
	for scanner.Scan() {
		line := scanner.Text()
		matched := true
		matched = matched && strings.HasPrefix(line, prefix) && strings.HasSuffix(line, suffix)
		for _, keyword := range keywords {
			matched = matched && strings.Contains(line, keyword)
		}
		if matched {
			fmt.Println(line)
			cnt++
			if nMax > 0 && cnt >= nMax {
				break
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
