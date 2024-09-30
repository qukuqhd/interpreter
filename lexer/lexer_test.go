package lexer

import (
	"testing"

	"github.com/qukuqhd/interpreter/token"
)

// 测试NextToken函数来完成词法分析器的工作
func Test_NextToken(t *testing.T) {
	input := `
	let five = 5;
	let ten = 10;
	let add = fn (x,y){
		x+y;
	};
	let result = add(five,ten);` //输入的代码
	tests := []struct { //希望得到的词法单元
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}
	l := NewLexer(input)
	for index, val := range tests {
		tok := l.NextToken()
		if tok.Type != val.expectedType {
			t.Fatalf("tests[%d] - 词法类型不一致 期望是 %q 实际上 %q", index, val.expectedType, tok.Type)
		}
		if tok.Literal != val.expectedLiteral {
			t.Fatalf("tests[%d] - 词法字面量不一致 期望是 %q 实际上 %q", index, val.expectedLiteral, tok.Literal)
		}
	}
}
