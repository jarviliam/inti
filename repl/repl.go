package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/jarviliam/inti/lexer"
	"github.com/jarviliam/inti/token"
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
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Fprintf(out, "%+v\n", tok)
		}
	}
}
