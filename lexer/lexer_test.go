package lexer

import (
  "testing"
  
  "csimple/token"
)

func TestLexer(t *testing.T) {
  inp := "run jump"
  tks := Lex(inp)
  
  res := []token.Token {
    {token.RunKw, "run", 0},
    {token.JumpKw, "jump", 0},
    {token.End, "", 0},
  }
  
  for i, tk := range tks {
    if tk.Type == token.Error {
      t.Fatalf("lexer error: %s", tk.Content)
    }
    
    if tk.Type != res[i].Type {
      t.Fatalf("type wrong. expected %s, got %s.", res[i].Type, tk.Type)
    }
    
    if tk.Content != res[i].Content {
      t.Fatalf("content wrong. expected %s, got %s. pos: %d", res[i].Content, tk.Content, tk.Pos)
    }
  }
}