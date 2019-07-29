package parser

import (
	"ast"
	"fmt"
	"lexer"
	"token"
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
		panic(t)
		return false
	}
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("Line %d : expected next token to be %s, got %s[%s] instead",
		p.l.LineN, t, p.peekToken.Type, p.peekToken.Literal)
	p.errors = append(p.errors, msg)
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Tables = []ast.Table{}

	table := p.parseTable()
	if table != nil {
		program.Tables = append(program.Tables, *table)
	}
	return program
}

func (p *Parser) parseTable() *ast.Table {
	table := &ast.Table{}
	if !p.expectPeek(token.IDENT) {
		return nil
	}
	table.Name = p.curToken.Literal
	p.nextToken()
	if p.curTokenIs(token.AS) {
		p.nextToken()
		p.nextToken()
	}
	if !p.curTokenIs(token.LBRACE) {
		return nil
	}
	p.nextToken()
	for {
		if !p.curTokenIs(token.ASTERISK) {
			break
		}
		column := p.parseColumn()
		if column == nil {
			return nil
		}
		table.Columns = append(table.Columns, *column)
	}
	for {
		if p.curTokenIs(token.PRIMARY_KEY) ||p.curTokenIs(token.KEY) || p.curTokenIs(token.UNIQUE_KEY) {
			if !p.handleKey(table) {
				return nil
			}
		}else {
			break
		}
	}
	//comment
	if p.curTokenIs(token.COMMENT) {
		if !p.expectPeek(token.IDENT) {
			return nil
		}
		table.Comment = p.curToken.Literal
	}
	return table
}

func (p *Parser) handleKey(table *ast.Table) bool {
	if p.curTokenIs(token.PRIMARY_KEY) {
		if !p.expectPeek(token.IDENT) {
			return false
		}
		table.Pk.Exist = true
		table.Pk.Keys = p.curToken.Literal
		p.nextToken()
	}
	if p.curTokenIs(token.KEY) {
		if !p.expectPeek(token.IDENT) {
			return false
		}
		table.Key.Exist = true
		table.Key.KeyName = p.curToken.Literal
		if !p.expectPeek(token.IDENT) {
			return false
		}
		table.Key.Keys = p.curToken.Literal
		p.nextToken()
	}
	if p.curTokenIs(token.UNIQUE_KEY) {
		if !p.expectPeek(token.IDENT) {
			return false
		}
		table.Uk.Exist = true
		table.Uk.KeyName = p.curToken.Literal
		if !p.expectPeek(token.IDENT) {
			return false
		}
		table.Uk.Keys = p.curToken.Literal
		p.nextToken()
	}
	return true
}

func (p *Parser) parseColumn() *ast.Column {
	column := &ast.Column{}
	//column的名字
	if !p.expectPeek(token.IDENT) {
		return nil
	}
	column.Name = p.curToken.Literal

	if !p.expectPeek(token.COLON) {
		return nil
	}

	if !p.expectPeek(token.IDENT) {
		return nil
	}

	column.Type = p.curToken.Literal

	p.nextToken()

	if p.curTokenIs(token.COMMENT) {
		p.nextToken()
		column.Comment = p.curToken.Literal
		p.nextToken()
	}
	return column
}
