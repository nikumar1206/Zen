package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"interpreter/lexer"
	"interpreter/repl"
)

var fileVar string

func init() {
	flag.StringVar(&fileVar, "f", "", "Determines to execute a file or initiate a repl")
	flag.Parse()
}

func main() {

	if fileVar == "" {
		err := repl.Initiate()
		check(err)
	} else {
		tokenizeFile()
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func tokenizeFile() {
	file, err := os.Open(fileVar)
	check(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		buff := scanner.Text()
		l := lexer.NewLexer(buff)

		tokens, err := l.Tokenize()

		check(err)
		for i := range tokens {
			fmt.Printf("Type: %s, Value: %s\n", tokens[i].Type, tokens[i].Value)
		}
	}
}
