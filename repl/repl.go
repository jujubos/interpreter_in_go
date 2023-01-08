package repl

import (
	"bufio"
	"fmt"
	"interpreter_in_go/lexer"
	"interpreter_in_go/token"
	"io"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	reader := bufio.NewReader(in)

	for {
		fmt.Fprintf(out, PROMPT)
		line, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
		}

		l := lexer.New(string(line))

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Fprintf(out, "%+v\n", tok)
		}
	}
}
