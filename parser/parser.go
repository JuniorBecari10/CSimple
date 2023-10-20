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

func (p *Parser) Parse() []ast.Statement {
  
}
