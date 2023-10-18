package lexer

import (
	"fmt"
	"strconv"
	"unicode"

	"csimple/token"
	"csimple/util"
)

type Lexer struct {
	input string
	data  util.FileData

	start  int
	cursor int

	startPos token.Position
	pos      token.Position
}

func New(input string, data util.FileData) Lexer {
	return Lexer{
		input: input,
		data:  data,

		start:  0,
		cursor: 0,

		startPos: token.Position{
			Line: 0,
			Col:  0,
		},

		pos: token.Position{
			Line: 0,
			Col:  0,
		},
	}
}

func (l *Lexer) Lex() ([]token.Token, bool) {
	tokens := []token.Token{}
	hadError := false

	for l.cursor < len(l.input) {
		l.skipWhitespace()

		l.start = l.cursor
		l.startPos = l.pos

		switch l.advance() {
		case '\n':
			tokens = append(tokens, l.newTokenAbs(token.NewLine, "", ""))
			//l.pos.Line++
			//l.pos.Col = 0

			fallthrough
		case '\t' | '\r':
			fallthrough
		case ' ':
			l.cursor++
			continue

		case '+':
			if l.matchAdvance('=') {
				tokens = append(tokens, l.newToken(token.PlusAssign))
				continue
			}

			tokens = append(tokens, l.newToken(token.Plus))

		case '-':
			if l.matchAdvance('=') {
				tokens = append(tokens, l.newToken(token.MinusAssign))
				continue
			}

			tokens = append(tokens, l.newToken(token.Minus))

		case '*':
			if l.matchAdvance('=') {
				tokens = append(tokens, l.newToken(token.TimesAssign))
				continue
			}

			tokens = append(tokens, l.newToken(token.Times))

		case '/':
			if l.matchAdvance('=') {
				tokens = append(tokens, l.newToken(token.DivideAssign))
				continue
			}

			tokens = append(tokens, l.newToken(token.Divide))

		case '%':
			if l.matchAdvance('=') {
				tokens = append(tokens, l.newToken(token.ModAssign))
				continue
			}

			tokens = append(tokens, l.newToken(token.Mod))

		case '^':
			if l.matchAdvance('=') {
				tokens = append(tokens, l.newToken(token.PowerAssign))
				continue
			}

			tokens = append(tokens, l.newToken(token.Power))

		case '&':
			if l.matchAdvance('=') {
				tokens = append(tokens, l.newToken(token.AndAssign))
				continue
			}

			tokens = append(tokens, l.newToken(token.And))

		case '|':
			if l.matchAdvance('=') {
				tokens = append(tokens, l.newToken(token.OrAssign))
				continue
			}

			tokens = append(tokens, l.newToken(token.Or))

		case '=':
			if l.matchAdvance('=') {
				tokens = append(tokens, l.newToken(token.Equals))
				continue
			}

			tokens = append(tokens, l.newToken(token.Assign))

		case '!':
			if l.matchAdvance('=') {
				tokens = append(tokens, l.newToken(token.Different))
				continue
			}

			tokens = append(tokens, l.newToken(token.Bang))

		case '<':
			if l.matchAdvance('=') {
				tokens = append(tokens, l.newToken(token.LessEq))
				continue
			}

			tokens = append(tokens, l.newToken(token.Less))

		case '>':
			if l.matchAdvance('=') {
				tokens = append(tokens, l.newToken(token.GreaterEq))
				continue
			}

			tokens = append(tokens, l.newToken(token.Greater))

		case '(':
			tokens = append(tokens, l.newToken(token.LParen))

		case ')':
			tokens = append(tokens, l.newToken(token.RParen))

		default:
			if l.isDigit(l.peek()) {
				l.advance()

				for l.isDigit(l.peek()) || l.peek() == '.' {
					l.advance()
				}

				n, err := strconv.ParseFloat(l.input[l.start:l.cursor], 64)

				if err != nil {
					util.ThrowError(l.data, l.startPos, fmt.Sprintf("Invalid number: '%s'", l.input[l.start:l.cursor]))
					hadError = true

					continue
				}

				tokens = append(tokens, l.newTokenLit(token.Number, n))
        continue
			}

			if l.isIdent(l.peek()) {
				l.advance()

				for l.isIdent(l.peek()) || l.isDigit(l.peek()) {
					l.advance()
				}

				tokens = append(tokens, l.newToken(token.Identifier))
				continue
			}

			// else
			l.recede()
			util.ThrowError(l.data, l.pos, fmt.Sprintf("Unknown token: '%c'", l.input[l.cursor]))
			l.advance()
			
			hadError = true
		}
	}

	return tokens, hadError
}

func (l *Lexer) skipWhitespace() {
	for unicode.IsSpace(rune(l.peek())) && l.peek() != '\n' {
		l.advance()
	}
}

func (l *Lexer) peek() uint8 {
	if l.cursor >= len(l.input) {
    return 0
  }

  return l.input[l.cursor]
}

func (l *Lexer) advance() uint8 {
	c := l.peek()

	if c == '\n' {
		l.pos.Line++
		l.pos.Col = 0
	} else {
		l.cursor++
		l.pos.Col++
	}

	return c
}

// todo! check if it's out of bounds and change Line if so
func (l *Lexer) recede() {
  if l.pos.Col == 0 {
		if l.pos.Line == 0 {
			panic("line cannot be lower than 0!")
		}

		l.pos.Line--
		l.pos.Col = len(l.data.Lines[l.pos.Line])
	} else {
		l.cursor--
		l.pos.Col--
	}
}

func (l *Lexer) matchAdvance(c uint8) bool {
	if l.peek() == c {
		l.advance()
		return true
	}

	return false
}

func (l *Lexer) isDigit(c uint8) bool {
	return c >= '0' && c <= '9'
}

func (l *Lexer) isIdent(c uint8) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || c == '_'
}

func (l *Lexer) newToken(ty token.TokenType) token.Token {
	return token.Token{
		Type:    ty,
		Pos:     l.startPos,
		Lexeme:  l.input[l.start:l.cursor],
		Literal: l.input[l.start:l.cursor],
	}
}

func (l *Lexer) newTokenLit(ty token.TokenType, lit any) token.Token {
	return token.Token{
		Type:    ty,
		Pos:     l.startPos,
		Lexeme:  l.input[l.start:l.cursor],
		Literal: lit,
	}
}

func (l *Lexer) newTokenAbs(ty token.TokenType, lex string, lit any) token.Token {
	return token.Token{
		Type:    ty,
		Pos:     l.startPos,
		Lexeme:  lex,
		Literal: lit,
	}
}
