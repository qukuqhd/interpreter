package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/qukuqhd/Interpreter/lexer"
	"github.com/qukuqhd/Interpreter/token"
)

const PROMPT = "~>" //提示符
const ICON = `            __,__
   .--.  .-"     "-.  .--.
  / .. \/  .-. .-.  \/ .. \
 | |  '|  /   Y   \  |'  | |
 | \   \  \ 0 | 0 /  /   / |
  \ '- ,\.-"""""""-./, -' /
   ''-' /_   ^ ^   _\ '-''
       |  \._   _./  |
       \   \ '~' /   /
        '._ '-=-' _.'
           '-----'
`

// Start repl启动，从read读取代码，输出结果给writer
func Start(in io.Reader, out io.Writer) {
	fmt.Fprint(out, ICON) //图标输出
	scanner := bufio.NewScanner(in)
	for {
		fmt.Fprint(out, PROMPT)   //提示符输出
		scanned := scanner.Scan() //等待输入
		if !scanned {
			return
		}
		line := scanner.Text()
		l := lexer.NewLexer(line)
		tok := l.NextToken()
		fmt.Fprintf(out, "%v\n", tok)
		for tok.Type != token.EOF {
			tok = l.NextToken()
			fmt.Fprintf(out, "%v\n", tok)
		}
	}
}
