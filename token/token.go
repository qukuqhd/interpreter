package token

type TokenType string

const ( //词法单元的类型

	//标识符
	IDENT = "IDENT"
	INT   = "INT"

	//运算符
	ASSIGN = "="
	PLUS   = "+"

	//分割符
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	//关键字
	FUNCTION = "FUNCTION"
	LET      = "LET"

	EOF     = "EOF"     //文件结束
	ILLEGAL = "ILLEGAL" //未知的词法单元
)

// Token 词法单元
type Token struct {
	Type    TokenType //词法单元的类型
	Literal string    //词法单元实际包含的信息
}

// keyWords 关键词容器
var keyWords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

// 检查标识符定义是否为关键字
func LookupIndent(ident string) TokenType {
	if val, ok := keyWords[ident]; ok { //关键词容器里面存在说明是关键词
		return val
	} else { //否则为定义
		return IDENT
	}
}
