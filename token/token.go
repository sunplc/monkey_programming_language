package token

const (
    ILLEGAL = "ILLEGAL"
    EOF     = "EOF"

    // 标识符+字面量
    IDENT = "IDENT" // add, foobar, x, y, ...
    INT   = "INT"   // 1343456
    STRING = "STRING"   // "foobar"

    // 运算符
    ASSIGN   = "="
    PLUS     = "+"
    MINUS    = "-"
    BANG     = "!"
    ASTERISK = "*"
    SLASH    = "/"

	LT = "<"
	GT = ">"

    // 分隔符
    COMMA     = ","
    SEMICOLON = ";"

    LPAREN = "("
    RPAREN = ")"
    LBRACE = "{"
    RBRACE = "}"

    LBRACKET = "["
    RBRACKET = "]"

    COLON = ":"

    // 关键字
    FUNCTION = "FUNCTION"
    LET      = "LET"
	TRUE     = "TRUE"
    FALSE    = "FALSE"
    IF       = "IF"
    ELSE     = "ELSE"
    RETURN   = "RETURN"

	EQ = "=="
	NOT_EQ = "!="

    MACRO = "MACRO"
)

type TokenType string

type Token struct {
	Type TokenType
	Literal string
}

var keywords = map[string]TokenType{
    "fn": FUNCTION,
    "let": LET,
    "true": TRUE,
	"false": FALSE,
	"if": IF,
	"else": ELSE,
	"return": RETURN,
    "macro": MACRO,
}

func LookupIdent(ident string) TokenType {
    if tokType, ok := keywords[ident]; ok {
        return tokType
    }
    return IDENT
}