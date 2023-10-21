package ast

const (
  InputNum = "num"
  InputStr = "str"
  InputBool = "bool"
)

// --- Statements

type Statement interface {
  stat()
}

// Syntax: <ident> = <expression>
type VarStat struct {
  Name  string
  Value Expression
}

// Syntax: <ident> +|-|*|/|^|&| (|) | = <expression>
type OperationStat struct {
  Name  string
  Value Expression
  Op    string
}

// Syntax: print <expression>
type PrintStat struct {
  BreakLine   bool
  Value  Expression
}

// Syntax: goto :<label>
type GotoStat struct {
  Label string
}

// Syntax: if <expression> goto :<label>
type IfStat struct {
  Exp Expression
  Label      string
}

type ExpStat struct {
  Exp Expression
}

type LabelStat struct {
  Name string
}

type ExitStat struct {
  Code Expression
}

type RetStat struct {}

func (v VarStat)        stat() {}
func (o OperationStat)  stat() {}
func (p PrintStat)      stat() {}
func (l LabelStat)      stat() {}
func (e ExitStat)       stat() {}
func (r RetStat)        stat() {}
func (g GotoStat)       stat() {}
func (i IfStat)         stat() {}
func (e ExpStat)        stat() {}

// --- Expressions

type Expression interface {
  exp()
}

type IdentifierExp struct {
  Value string
}

type NumberExp struct {
  Value float64
}

type StringExp struct {
  Value string
}

type BoolExp struct {
  Value bool
}

type BinExp struct {
  NodeA Expression
  NodeB Expression
  Op string
}

type MinusExp struct {
  Value Expression
}

type InputExp struct {
  Type string
}

type FactorialExp struct {
  Exp Expression
}

type ExecExp struct {
  Command Expression
}

func (i IdentifierExp) exp() {}
func (n NumberExp)     exp() {}
func (s StringExp)     exp() {}
func (b BoolExp)       exp() {}
func (b BinExp)        exp() {}
func (m MinusExp)      exp() {}
func (i InputExp)      exp() {}
func (f FactorialExp)  exp() {}
func (e ExecExp)       exp() {}
