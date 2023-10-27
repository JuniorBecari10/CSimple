package parser

import (
	"csimple/ast"
	"csimple/token"
	"csimple/util"
	"fmt"
	"strconv"
)

type Parser struct {
	input  []token.Token
	cursor int
	data   util.FileData
}

func New(input []token.Token, data util.FileData) Parser {
	return Parser{
		input:  input,
		cursor: 0,
		data:   data,
	}
}

func (p *Parser) Parse() ([]ast.Statement, bool) {
	stats := []ast.Statement{}
	hadError := false

	for p.cursor < len(p.input) {
		for p.match(token.NewLine) {
			p.advance()
		}

		stat := p.statement()

		// an error occurred
		if stat == nil {
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
	next := p.peek(0)

	if len(p.input) >= 1 {
		switch tk.Type {
		case token.IfKw:
			return p.parseIfStat()

		case token.GotoKw:
			return p.parseGotoStat()

		case token.PrintKw:
			fallthrough
		case token.PrintlnKw:
			return p.parsePrintStat(tk)

		case token.RetKw:
			return ast.RetStat{}

		case token.ExitKw:
			return p.parseExitStat()
		}
	}

	if len(p.input) >= 2 {
		switch tk.Type {
		case token.Identifier:
			if next.Type == token.Assign {
				return p.parseVarStat(tk)
			}

			if Find(string(next.Type), []string { token.PlusAssign, token.MinusAssign, token.TimesAssign, token.DivideAssign, token.PowerAssign, token.ModAssign, token.AndAssign, token.OrAssign }) != -1 {
				return p.parseOperationStat(tk)
			}

		case token.Colon:
			return p.parseLabelStat()
		}
	}

	p.recede()
	exp := p.parseExp(0)

	if exp == nil {
		return nil
	}

	return ast.ExpStat{Exp: exp}
}

// ---

func (p *Parser) parseIfStat() ast.Statement {
	cond := p.parseExp(0)

	if cond == nil {
		return nil
	}

	if !p.matchAdvance(token.GotoKw) {
		util.ThrowError(p.data, p.peek(0).Pos, fmt.Sprintf("Expected 'goto' keyword after if condition, got '%s'.", p.peek(0).Lexeme))
		return nil
	}

	label := p.parseLabel()

	if label == "" {
		return nil
	}

	return ast.IfStat{
		Cond:  cond,
		Label: label,
	}
}

func (p *Parser) parseGotoStat() ast.Statement {
	label := p.parseLabel()

	if label == "" {
		return nil
	}

	return ast.GotoStat{
		Label: label,
	}
}

func (p *Parser) parsePrintStat(tk token.Token) ast.Statement {
	exp := p.parseExp(0)

	if exp == nil {
		return nil
	}

	return ast.PrintStat{
		BreakLine: tk.Type == token.PrintlnKw,
		Value:     exp,
	}
}

func (p *Parser) parseExitStat() ast.Statement {
	exp := p.parseExp(0)

	if exp == nil {
		return nil
	}

	return ast.ExitStat{
		Code: exp,
	}
}

/*
if tk.Type != token.Identifier {
  util.ThrowError(p.data, p.peek().Pos, "Variable names must be identifiers (e.g. only letters, numbers and underscores; cannot start with a number).")
  return nil
}
*/

func (p *Parser) parseVarStat(tk token.Token) ast.Statement {
	//fmt.Println(p.peek(0))
	if !p.matchAdvance(token.Assign) {
		util.ThrowError(p.data, p.peek(0).Pos, fmt.Sprintf("Expected '=' after variable name, got '%s'", p.peek(0).Lexeme))
		return nil
	}

	//fmt.Println(p.peek(0))
	//fmt.Println(p.advance())
	p.advance()
	exp := p.parseExp(0)

	if exp == nil {
		return nil
	}

	return ast.VarStat{
		Name:  tk.Lexeme,
		Value: exp,
	}
}

func (p *Parser) parseOperationStat(tk token.Token) ast.Statement {
	ope := p.advance().Lexeme

	if ope == "" {
		util.ThrowError(p.data, p.peek(-1).Pos, fmt.Sprintf("Expected operation, got '%s'.", p.peek(-1).Lexeme))
		return nil
	}

	exp := p.parseExp(0)

	if exp == nil {
		return nil
	}

	return ast.OperationStat{
		Name:  tk.Lexeme,
		Ope:   ope,
		Value: exp,
	}
}

func (p *Parser) parseLabelStat() ast.Statement {
	label := p.parseLabel()

	if label == "" {
		return nil
	}

	return ast.LabelStat{
		Name: label,
	}
}

// ---

func (p *Parser) parseLabel() string {
	if !p.matchAdvance(token.Colon) {
		util.ThrowError(p.data, p.peek(0).Pos, "Expected ':' before label name.")
		return ""
	}

	tk := p.advance()

	if tk.Type != token.Identifier {
		util.ThrowError(p.data, tk.Pos, "Expected label name (identifier) after ':'.")
	}

	return tk.Lexeme
}

// ---

/*
Precedence Order:

0 - And / Or
1 - Equality
2 - Less / Greater
3 - Sum / Sub
4 - Mul / Div
5 - Mod / Power
6 - Prefix
7 - Postfix
8 - Grouped (Parentheses)
9 - Primary
*/
func (p *Parser) parseExp(precedence int) ast.Expression {
	switch precedence {
	case 0:
		return p.bin(precedence, token.And, token.Or)

	case 1:
		return p.bin(precedence, token.Equals, token.Different)

	case 2:
		return p.bin(precedence, token.Less, token.LessEq, token.Greater, token.GreaterEq)

	case 3:
		return p.bin(precedence, token.Plus, token.Minus)

	case 4:
		return p.bin(precedence, token.Times, token.Divide)

	case 5:
		return p.bin(precedence, token.Mod, token.Power)

	case 6:
		if Find(p.peek(0).Lexeme, []string { "!", "-" }) != -1 {
			ope := p.advance().Lexeme
			exp := p.parseExp(precedence + 1)

			if exp == nil {
				return nil
			}

			return ast.UnaryExp{
				Exp: exp,
				Ope: ope,
			}
		}

		fallthrough

	case 7:
		exp := p.parseExp(precedence + 1)

		if exp == nil {
			return nil
		}	

		if Find(p.peek(0).Lexeme, []string { "!" }) != -1 {
			ope := p.advance().Lexeme

			return ast.UnaryExp{
				Exp: exp,
				Ope: ope,
			}
		}

		return exp

	case 8:
		if p.matchAdvance(token.LParen) {
			exp := p.parseExp(0)

			if exp == nil {
				return nil
			}

			if !p.matchAdvance(token.RParen) {
				util.ThrowError(p.data, p.peek(0).Pos, fmt.Sprintf("Expected ')' after expression, got '%s'.", p.peek(0).Lexeme))
				return nil
			}

			return exp
		}

	case 9:
		tk := p.advance()
		fmt.Println(tk.Type)

		switch tk.Type {
		case token.Identifier:
			return ast.IdentifierExp{
				Value: tk.Lexeme,
			}

		case token.Number:
			fmt.Println("number!")
			value, err := strconv.ParseFloat(tk.Lexeme, 64)

			if err != nil {
				util.ThrowError(p.data, tk.Pos, fmt.Sprintf("Invalid number: '%s'", tk.Lexeme))
				return nil
			}

			return ast.NumberExp{
				Value: value,
			}

		case token.TrueKw:
			fallthrough
		case token.FalseKw:
			return ast.BoolExp{
				Value: tk.Type == token.TrueKw,
			}

		case token.InputKw:
			ty := ""

			switch p.peek(0).Type {
			case token.TypeNum:
				ty = ast.InputNum

			case token.TypeStr:
				ty = ast.InputStr

			case token.TypeBool:
				ty = ast.InputBool
			}

			return ast.InputExp{
				Type: ty,
			}

		case token.ExecKw:
			exp := p.parseExp(0)

			return ast.ExecExp{
				Command: exp,
			}
		}
	}

	util.ThrowError(p.data, p.peek(-1).Pos, fmt.Sprintf("Expected expression, got '%s'.", p.peek(-1).Lexeme))
	return nil
}

func (p *Parser) bin(precedence int, types ...token.TokenType) ast.Expression {
	left := p.parseExp(precedence + 1)

	if left == nil {
		return nil
	}

	ope := ""

	for _, tk := range types {
		if p.matchAdvance(tk) {
			ope = p.peek(0).Lexeme
			break
		}
	}

	if ope == "" {
		util.ThrowError(p.data, p.peek(0).Pos, fmt.Sprintf("Expected an infix operator, got '%s'.", p.peek(0).Lexeme))
		return nil
	}

	right := p.parseExp(precedence + 1)

	if right == nil {
		return nil
	} else {
		left = ast.BinExp{
			Left:  left,
			Right: right,
			Ope:   ope,
		}
	}

	return left
}

// ---

func (p *Parser) peek(offset int) token.Token {
	if p.cursor + offset >= len(p.input) {
		return token.Token{
			Lexeme: "",
			Pos: token.Position{
				Line: p.input[p.cursor - 1 + offset].Pos.Line,
				Col:  p.input[p.cursor - 1 + offset].Pos.Col + len(p.input[p.cursor-1].Lexeme),
			},
		}
	}

	return p.input[p.cursor + offset]
}

func (p *Parser) match(t token.TokenType) bool {
	return p.peek(0).Type == t
}

func (p *Parser) advance() token.Token {
	t := p.peek(0)
	p.cursor++

	return t
}

func (p *Parser) recede() {
	p.cursor--
}

func (p *Parser) matchAdvance(t token.TokenType) bool {
	if p.match(t) {
		p.advance()
		return true
	}

	return false
}

// ---

func Find(what string, where []string) int {
	for i, v := range where {
		if v == what {
			return i
		}
	}

	return -1
}
