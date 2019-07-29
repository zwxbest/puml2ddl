package token

type TokenType string

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENT = "IDENT" // add, foobar, x, y, ...

	COLON = ":"
	LINE = "--"
	SEMICOLON = ";"
	ASTERISK ="*"

	QUOTATION = "\""

	LBRACE = "{"
	RBRACE = "}"

	SKINPARAM = "skinparam"
	COLUMN = "column"
	COLUMN_TYPE = "column_type"


	// Keywords
	AS = "as"
	ENTITY = "ENTITY"
	COMMENT = "comment"
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
	"comment":     COMMENT,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
