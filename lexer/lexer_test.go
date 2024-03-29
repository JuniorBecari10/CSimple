package lexer

import (
	"strings"
	"testing"

	"csimple/token"
	"csimple/util"
)

func TestLexer(t *testing.T) {
	input := "x = 10"
	l := New(input, util.FileData{
		Name:  "test.go",
		Lines: strings.Split(input, "\n"),
	})
	tokens, hadError := l.Lex()

	if hadError {
		t.Errorf("Lexer error\n")
	}

	expected := []token.Token{
		{
			Type: token.Identifier,
			Pos: token.Position{
				Line: 0,
				Col:  0,
			},
			Lexeme:  "x",
			Literal: "x",
		},

		{
			Type: token.Assign,
			Pos: token.Position{
				Line: 0,
				Col:  2,
			},
			Lexeme:  "=",
			Literal: "=",
		},

		{
			Type: token.Number,
			Pos: token.Position{
				Line: 0,
				Col:  4,
			},
			Lexeme:  "10",
			Literal: 10.0,
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
