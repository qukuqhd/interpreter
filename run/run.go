package run

import (
	"fmt"
	"io"
	"log"

	"github.com/qukuqhd/Interpreter/lexer"
	"github.com/qukuqhd/Interpreter/token"
)

func Run(reader io.Reader, output io.Writer) {
	script, err := io.ReadAll(reader)
	if err != nil {
		log.Fatal(err)
	} else {
		l := lexer.NewLexer(string(script))
		for {
			tok := l.NextToken()
			fmt.Fprintf(output, "%v\n", tok)
			if tok.Type == token.EOF {
				break
			}
		}
	}
}
