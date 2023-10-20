package parser

import (
	"csimple/ast"
	"csimple/token"
	"csimple/util"
)

type Parser struct {
  input  []token.Token
  cursor int
  data   util.FileData
}

func New(input []token.Token, data util.FileData) Parser {
  return Parser{
    input: input,
    cursor: 0,
    data: data,
  }
}

func (p *Parser) Parse() ([]ast.Statement, bool) {
  stats := []ast.Statement {}
  hadError := false

  for p.cursor < len(p.input) {
    for p.peek().Lexeme != "" && p.match(token.NewLine) {
      p.advance()
    }

    if p.peek().Lexeme != "" {
      break
    }

    ret := nil

    
  }

  return stats, hadError
}

func (p *Parser) peek() token.Token {
  if p.cursor >= len(p.input) {
    return token.Token{Lexeme: ""}
  }

  return p.input[p.cursor]
}

func (p *Parser) match(t token.TokenType) bool {
  return p.peek().Type == t
}

func (p *Parser) advance() token.Token {
  t := p.peek()
  p.cursor++

  return t
}
