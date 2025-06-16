package main

import (
	"bufio"
	"fmt"
	"io"
)

const PROMPT = "> "

func StartREPL(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	for {
		fmt.Println(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := NewLexer(line)
		for word := l.NextWord(); word.Type != EOF; word = l.NextWord() {
			fmt.Printf("%+v\n", word)
		}
	}
}
