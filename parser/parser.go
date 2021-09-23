package parser

import (
	"github.com/jarviliam/inti/ast"
	"github.com/jarviliam/inti/lexer"
	"github.com/jarviliam/inti/token"
)

type Parser struct {
	l       *lexer.Lexer
	currTok token.Token
	peekTok token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}
	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) nextToken() {
	p.currTok = p.peekTok
	p.peekTok = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	return nil
}
