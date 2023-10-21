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

    stat := p.statement()

    // an error occurred; either 'stat' is nil or it is an ExpStat and its Exp field is nil
    if exp, ok := stat.(ast.ExpStat); (ok && exp.Exp == nil) || stat == nil {
      hadError = true

      for p.cursor < len(p.input) && p.match(token.NewLine) {
        p.advance()
      }

      continue
    }

    stats = append(stats, stat)
  }

  return stats, hadError
}

// ---

func (p *Parser) statement() ast.Statement {
  tk := p.advance()

  if len(p.input) >= 1 {
    switch tk.Type {
    case token.IfKw:
      return p.parseIfStat()

    case token.GotoKw:
      return p.parseGotoStat()

    case token.PrintKw:
      fallthrough
    case token.PrintlnKw:
      return p.parsePrintStat()

    case token.RetKw:
      return p.parseRetStat()

    case token.ExitKw:
      return p.parseExitStat()
    }
  }

  return ast.ExpStat { Exp: p.parseExp() }
}

// ---

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
