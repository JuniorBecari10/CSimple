package parser

import (
	"csimple/ast"
	"csimple/lexer"
	"csimple/util"
	"reflect"
	"strings"
	"testing"
)

func TestGenerateTest(t *testing.T) {
	input := "x = 10"
	data := util.FileData{
		Name:  "test!",
		Lines: strings.Split(input, "\n"),
	}

	l := lexer.New(input, data)
	tokens, hadError := l.Lex()

	if hadError {
		t.Errorf("Lexer error.\n")
	}
	
	parser := New(tokens, data)
	stats, hadError := parser.Parse()

	if hadError {
		t.Errorf("Parser error.\n")
	}

  expect := []ast.Statement {
    ast.VarStat {
      Name: "x",
      Value: ast.NumberExp {
        Value: 10.0,
      },
    },
  }

  if !reflect.DeepEqual(stats, expect) {
    t.Errorf("Expected %v, but got %v", expect, stats)
  }
}
