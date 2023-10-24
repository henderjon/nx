package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n int
	flag.Func("n", "the number of times to execute", func(s string) error {
		n = intval(s)
		return nil
	})
	flag.Parse()

	args := flag.Args()

	if n <= 0 {
		n = intval(flag.Arg(0))
		args = args[1:]
	}

	cmd := strings.Join(args, " ")

	for i := 0; i < n; i++ {
		fmt.Println(cmd)
	}
}

func intval(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	return n
}
