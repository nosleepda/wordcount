package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	words := scanner.Text()

	fmt.Println(len(strings.Split(words, " ")))
}
