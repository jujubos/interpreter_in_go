package repl

import (
	"bufio"
	"fmt"
	"interpreter_in_go/lexer"
	"interpreter_in_go/parser"
	"io"
)

const PROMPT = ">> "

//func Start(in io.Reader, out io.Writer) {
//	reader := bufio.NewReader(in)
//
//	for {
//		fmt.Fprintf(out, PROMPT)
//		line, _, err := reader.ReadLine()
//		if err != nil {
//			if err == io.EOF {
//				break
//			}
//		}
//
//		l := lexer.New(string(line))
//
//		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
//			fmt.Fprintf(out, "%+v\n", tok)
//		}
//	}
//}

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		io.WriteString(out, program.String())
		io.WriteString(out, "\n")
	}
}

func printParserErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
