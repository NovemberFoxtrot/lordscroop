package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strings"
)

var workers = runtime.NumCPU()

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		somestring := scanner.Text()

		if strings.Contains(somestring, "*") {
			fmt.Println(somestring)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
