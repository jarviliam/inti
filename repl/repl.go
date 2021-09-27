package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/jarviliam/inti/lexer"
	"github.com/jarviliam/inti/parser"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	s := bufio.NewScanner(in)

	for {
		fmt.Fprintf(out, PROMPT)
		scanned := s.Scan()
		if !scanned {
			return
		}
		line := s.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserError(out, p.Errors())
			continue
		}
		io.WriteString(out, program.String())
		io.WriteString(out, "\n")
	}
}

func printParserError(out io.Writer, errors []string) {
	for _, m := range errors {
		io.WriteString(out, "\t"+m+"\n")
	}
}
