package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var x bool
	flag.BoolFunc("x", "read the remainder of the args as the command", func(s string) error {
		x = true
		return nil
	})

	var n int
	flag.Func("n", "the number of times to execute; if absent, use the first arg", func(s string) error {
		n = intval(s)
		return nil
	})
	flag.Parse()

	args := flag.Args()

	if n <= 0 {
		n = intval(flag.Arg(0))
		args = args[1:]
	}

	var cmd string

	if x {
		cmd = strings.Join(args, " ")
	} else {
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			if err := scanner.Err(); err != nil {
				fmt.Fprintln(os.Stderr, "unable to read from stdin")
			}
			cmd = scanner.Text()
		}
	}

	for i := 0; i < n; i++ {
		fmt.Println(cmd)
	}
}

func intval(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to parse '%s' as a number\n", s)
		os.Exit(1)
	}
	return n
}
