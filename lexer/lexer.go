package lexer

import "github.com/qukuqhd/Interpreter/token"

// Lexer 词法分析器
type Lexer struct {
	input        []rune //输入的代码
	position     int    //词法分析到的位置
	readPosition int    //词法分析到要读取的位置
	ch           rune   //当前读取到的unicode字符
}

// NewLexer 传入代码得到词法分析器
func NewLexer(input string) *Lexer {
	l := &Lexer{
		input: []rune(input),
	}
	l.readChar()
	return l
}

// 读取字节
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) { //下一个位置已经超出了
		l.ch = 0 //读取到的字节就是0终止
	} else { //否则读取下一个字节
		l.ch = l.input[l.readPosition]
	}
	//位置往后移动
	l.position = l.readPosition
	l.readPosition++
}

func newToken(tokenType token.TokenType, ch rune) *token.Token {
	return &token.Token{
		Type:    tokenType,
		Literal: string(ch),
	}
}

// NextToken 分析得到词法单元
func (l *Lexer) NextToken() *token.Token {
	var tok = &token.Token{}

	//对缩进空格换行进行跳过
	l.skipSpace()
	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar() //往下一个读取字符
			tok.Type = token.EQ
			tok.Literal = string(ch) + string(l.ch)
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '/':
		tok = newToken(token.SLASH, l.ch)
	case '>':
		tok = newToken(token.GT, l.ch)
	case '<':
		tok = newToken(token.LT, l.ch)
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar() //往下一个读取字符
			tok.Type = token.NOT_EQ
			tok.Literal = string(ch) + string(l.ch)
		} else {
			tok = newToken(token.BANG, l.ch)
		}
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case 0:
		tok = &token.Token{
			Type:    token.EOF,
			Literal: "",
		}
	default: //检测到的其他字符的处理
		if isLetter(l.ch) { //是合法的标识符字符开头才读取标识符
			tok.Literal = l.readIdentifier()           //读取字符串
			tok.Type = token.LookupIndent(tok.Literal) //分析字符串是标识符还是应该是定义
			return tok                                 //这里直接return 是以为字符串截取读取就读取到了下一个词法的位置去了
		} else if isDight(l.ch) { //是否是数字
			//读取整个数字
			tok.Literal = l.readNumber()
			tok.Type = token.INT
			return tok //这里直接return 是以为字符串截取读取就读取到了下一个词法的位置去了
		} else { //否则为未知词法单元
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar() //往下读取分析
	return tok
}

// readIdentifier 读取标识符符
func (l *Lexer) readIdentifier() string {
	begin, end := l.sliceByRule(isLetter)
	return string(l.input[begin:end])
}

// isLetter 判断是否为合法的标识符的字符
func isLetter(ch rune) bool {
	return !(ch == 0 || ('0' <= ch && ch <= '9') || ch == ' ' || ch == '\r' || ch == '\n' || ch == '\t' || ch == '(' || ch == ')' || ch == '{' || ch == '}' || ch == '+' || ch == '-' || ch == '*' || ch == '/' || ch == '\'' || ch == '\\' || ch == ',' || ch == '.' || ch == '^' || ch == ';')
}

// skipSpace 跳过不必的换行空白
func (l *Lexer) skipSpace() {
	for l.ch == ' ' || l.ch == '\n' || l.ch == '\r' || l.ch == '\t' {
		l.readChar() //往下走接着读取
	}
}

// isDight 判断是否为数字
func isDight(ch rune) bool {
	return '0' <= ch && ch <= '9'
}

// 读取开始的数字
func (l *Lexer) readNumber() string {
	begin, end := l.sliceByRule(isDight)
	return string(l.input[begin:end])
}

//! sliceByRule 根据规则截取到范围 具有副作用会修改position
func (l *Lexer) sliceByRule(rule func(rune) bool) (begin, end int) {
	begin = l.position
	for rule(l.ch) {
		l.readChar()
	}
	end = l.position
	return
}

//查看下一个字符的内容 peekChar
func (l *Lexer) peekChar() rune {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}
