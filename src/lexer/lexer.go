package lexer

import (
	"token"
)

type Lexer struct {
	LineN        int
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
}

func New(input string, lineN int) *Lexer {
	l := &Lexer{input: input, LineN: lineN}
	l.readChar()
	return l
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.nextLine()
	l.skip()

	switch l.ch {
	case ':':
		tok = newToken(token.COLON, l.ch)
	case '"':
		l.readChar()
		tok.Literal = l.readQuote()
		tok.Type = token.IDENT
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isIdent(l.ch) {
			tok.Literal = l.readIdent()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		}
	}

	l.readChar()
	return tok
}

func (l *Lexer) nextLine()  {
	if l.ch == '\n' {
		l.LineN ++
		l.readChar()
	}
}

func (l *Lexer) skip() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\r' || l.ch == '-'  {
		l.readChar()
	}
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func (l *Lexer) readQuote() string {
	position := l.position
	for {
		if l.ch == '"'{
			break
		}
		l.readChar()
	}
	return l.input[position:l.position]
}


func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) || isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_' || ch == '(' || ch == ')'
}

func isIdent(ch byte) bool {
	return ch != '*' && ch != ':' && ch != '"' && ch != '{' && ch != '}' && ch != 0 && ch != ' ' && ch != '\t' && ch != '\r' && ch != '\n'
}

func (l *Lexer) readIdent() string {
	position := l.position
	for isIdent(l.ch){
		l.readChar()
	}
	return l.input[position:l.position]
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
