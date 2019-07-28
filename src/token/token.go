package token

type TokenType string

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENT = "IDENT" // add, foobar, x, y, ...

	COLON = ":"
	LINE = "--"
	COMMA     = ","
	SEMICOLON = ";"
	ASTERISK ="*"
	LL ="<<"
	RR =">>"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	SKINPARAM = "skinparam"
	COLUMN = "column"
	COLUMN_TYPE = "column_type"


	// Keywords
	AS = "as"
	ENTITY = "ENTITY"
	COMMENT = "commnet"
	KEY = "key"
	PRIMARY_KEY = "primary_key"
	UNIQUE_KEY  = "unique_key"
)

type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"as":     AS,
	"entity":    ENTITY,
	"<<KEY>>":   KEY,
	"<<PK>>":  PRIMARY_KEY,
	"<<UK>>":     UNIQUE_KEY,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT;
}
