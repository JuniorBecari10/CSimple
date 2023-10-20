package ast

const (
  InputNum = "num"
  InputStr = "str"
  InputBool = "bool"
)

type Statement interface {
  stat()
}

// Syntax: <ident> = <expression>
type Var struct {
  Name  string
  Value Expression
}

// Syntax: <ident> +|-|*|/|^|&| (|) | = <expression>
type Operation struct {
  Name  string
  Value Expression
  Op    string
}

// Syntax: print <expression>
type Print struct {
  BreakLine   bool
  Value  Expression
}

// Syntax: goto :<label>
type Goto struct {
  Label string
}

// Syntax: if <expression> goto :<label>
type If struct {
  Exp Expression
  Label      string
}

type Exp struct {
  Exp Expression
}

type Label struct {
  Name string
}

type Exit struct {
  Code Expression
}

type Ret struct {}

func (v Var)        stat() {}
func (o Operation)  stat() {}
func (p Print)      stat() {}
func (l Label)      stat() {}
func (e Exit)       stat() {}
func (r Ret)        stat() {}
func (g Goto)       stat() {}
func (i If)         stat() {}
func (e Exp)        stat() {}

// Expressions

type Expression interface {
  exp()
}

type Identifier struct {
  Value string
}

type Number struct {
  Value float64
}

type String struct {
  Value string
}

type Bin struct {
  NodeA Expression
  NodeB Expression
  Op string
}

type Minus struct {
  Value Expression
}

type Input struct {
  Type string
}

type Bool struct {
  Value bool
}

type Factorial struct {
  Exp Expression
}

type Exec struct {
  Command Expression
}

func (i Identifier) exp() {}
func (n Number)     exp() {}
func (s String)     exp() {}
func (b Bin)        exp() {}
func (m Minus)      exp() {}
func (i Input)      exp() {}
func (b Bool)       exp() {}
func (f Factorial)  exp() {}
func (e Exec)       exp() {}
