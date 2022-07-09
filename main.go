package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	flag.Parse()
	words := strings.Join(flag.Args(), "")

	if len(words) == 0 {
		fmt.Println(0)

		return
	}

	fmt.Println(len(strings.Split(words, " ")))
}

//func readInput() (pattern, words string, err error) {
//	flag.StringVar(&pattern, "p", "", "pattern to match against")
//	flag.Parse()
//	if pattern == "" {
//		return pattern, src, errors.New("missing pattern")
//	}
//	src = strings.Join(flag.Args(), "")
//	if src == "" {
//		return pattern, src, errors.New("missing string to match")
//	}
//	return pattern, src, nil
//}
