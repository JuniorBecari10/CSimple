package lexer

import (
	"testing"

	"csimple/token"
	"csimple/util"
)

func TestLexer(t *testing.T) {
  input := "123.4 hello_"
	l := New(input, util.FileData {
    Name: "test.go",
    Lines: []string { input },
  })
	tokens := l.Lex()

	expected := []token.Token{
		{
			Type: token.Number,
			Pos: token.Position{
				Line: 0,
				Col:  0,
			},
			Lexeme:  "123.4",
			Literal: 123.4,
		},
    
		{
			Type: token.Identifier,
			Pos: token.Position{
				Line: 0,
				Col:  6,
			},
			Lexeme:  "hello_",
			Literal: "hello_",
		},
	}
	if len(tokens) != len(expected) {
		t.Errorf("Expected %d tokens, but got %d", len(expected), len(tokens))
	}

	for i := 0; i < len(tokens); i++ {
		if tokens[i].Type != expected[i].Type {
			t.Errorf("Expected token type %v, but got %v", expected[i].Type, tokens[i].Type)
		}

		if tokens[i].Pos.Line != expected[i].Pos.Line {
			t.Errorf("Expected token line %d, but got %d", expected[i].Pos.Line, tokens[i].Pos.Line)
		}

		if tokens[i].Pos.Col != expected[i].Pos.Col {
			t.Errorf("Expected token column %d, but got %d", expected[i].Pos.Col, tokens[i].Pos.Col)
		}

		if tokens[i].Lexeme != expected[i].Lexeme {
			t.Errorf("Expected token lexeme '%s', but got '%s'", expected[i].Lexeme, tokens[i].Lexeme)
		}

		if tokens[i].Literal != expected[i].Literal {
			t.Errorf("Expected token literal '%s', but got '%s'", expected[i].Literal, tokens[i].Literal)
		}
	}
}
