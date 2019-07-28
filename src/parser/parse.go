package parser

import (
	"lexer"
	"token"
	"fmt"
	"ast"
)

type Parser struct {
	l      *lexer.Lexer
	errors []string

	curToken  token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:      l,
		errors: []string{},
	}

	// Read two tokens, so curToken and peekToken are both set
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		p.peekError(t)
		return false
	}
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead",
		t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Tables = []ast.Table{}

	for !p.curTokenIs(token.EOF) {
		table := p.parseTable()
		if table != nil {
			program.Tables = append(program.Tables, *table)
		}
		p.nextToken()
	}
	return program
}

func (p *Parser) parseTable() *ast.Table {
	table := &ast.Table{}

	//table的名字
	if !p.expectPeek(token.IDENT) {
		return nil
	}

	table.Name = p.curToken.Literal

	if p.peekTokenIs(token.AS) {
		p.nextToken()
		p.nextToken()
	}

	if p.peekTokenIs(token.LBRACE) {
		p.nextToken()
	}

	if p.peekTokenIs(token.ASTERISK) {

	}

	table.Columns = append(table.Columns, *p.parseColumn())

	stmt.Value = p.parseExpression(LOWEST)

	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseColumn() *ast.Column {
	column := &ast.Column{}
	//column的名字
	if !p.expectPeek(token.IDENT) {
		return nil
	}

	column.Name = p.curToken.Literal

	if p.peekTokenIs(token.COLON) {
		p.nextToken()
	}

	if p.peekTokenIs(token.LBRACE) {
		p.nextToken()
	}

	if p.peekTokenIs(token.ASTERISK) {

	}
	p.nextToken()

	stmt.Value = p.parseExpression(LOWEST)

	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}
