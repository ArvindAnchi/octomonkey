package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENT = "IDENT"
	INT   = "INT"

	ASSIGN = "="
	PLUS   = "+"
	MINUS  = "-"
	STAR   = "*"
	SLASH  = "/"

	EQ   = "=="
	N_EQ = "!="

	BANG      = "!"
	COMMA     = ","
	SEMICOLON = ";"

	L_THAN = "<"
	G_THAN = ">"

	L_PAREN = "("
	R_PAREN = ")"
	L_BRACE = "{"
	R_BRACE = "}"

	FUNCTION = "FUNCTION"
	LET      = "LET"

	IF     = "IF"
	THEN   = "THEN"
	ELSE   = "ELSE"
	TRUE   = "TRUE"
	FALSE  = "FALSE"
	RETURN = "RETURN"
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"if":     IF,
	"then":   THEN,
	"else":   ELSE,
	"true":   TRUE,
	"false":  FALSE,
	"return": RETURN,
}

func LookupIdent(ident string) TokenType {
	if token, ok := keywords[ident]; ok {
		return token
	}

	return IDENT
}
