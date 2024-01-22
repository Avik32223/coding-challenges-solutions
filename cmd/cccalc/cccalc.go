package main

import (
	"bufio"
	"flag"
	"io"
	"os"
	"strings"

	"github.com/Avik32223/coding-challenges-solutions/internal/calculator"
)

func main() {
	flag.Parse()
	args := flag.Args()
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	if len(args) == 1 {
		reader = bufio.NewReader(strings.NewReader(args[0]))
	}
	source, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	result, err := calculator.Calculate(source)
	if err != nil {
		panic(err)
	}
	writer.WriteString(result)
	writer.WriteRune('\n')
	writer.Flush()
}
