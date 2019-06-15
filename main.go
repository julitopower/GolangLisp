package main

import "fmt"

////////////////////////////////////////////////////////////////////////////////
// Literal
////////////////////////////////////////////////////////////////////////////////
type Int32Lit struct {
	value int
}
type I32 = Int32Lit

func (_ Int32Lit) isAtom()            {}
func (_ Int32Lit) isSExpr()           {}
func (v Int32Lit) Value() interface{} { return v.value }

type Float32Lit struct {
	value float32
}
type F32 = Float32Lit

func (_ Float32Lit) isAtom()            {}
func (_ Float32Lit) isSExpr()           {}
func (v Float32Lit) Value() interface{} { return v.value }

////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////
// Symbol
////////////////////////////////////////////////////////////////////////////////
const (
	SymVar = iota
	SymProc
	SymUnbound
)

type Symbol struct {
	Name  string
	Type  int
	value interface{}
}

func S(name string) Symbol {
	return Symbol{name, SymUnbound, nil}
}

func V(name string, value interface{}) Symbol {
	return Symbol{name, SymVar, value}
}

func NewSymbol(name string) Symbol {
	return Symbol{Name: name}
}

func (s Symbol) String() string {
	return fmt.Sprintf(s.Name)
}

func (_ Symbol) isAtom()  {}
func (_ Symbol) isSExpr() {}
func (s Symbol) Value() interface{} {
	return s.value
}

////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////
// SExprList
////////////////////////////////////////////////////////////////////////////////
type SExprList struct {
	Expr []SExpr
}

func NewList(data []SExpr) SExprList {
	return SExprList{data}
}

func (_ SExprList) isList()  {}
func (_ SExprList) isSExpr() {}

// Add an SExpr to a list
func (sexpr *SExprList) Add(data SExpr) {
	sexpr.Expr = append(sexpr.Expr, data)
}

func (sexpr SExprList) Value() interface{} {
	return sexpr
}

////////////////////////////////////////////////////////////////////////////////

// This is simply a collection of SExpr
type List interface {
	isList()
	SExpr
}

// Atoms can be symbols or literals
type Atom interface {
	isAtom()
	SExpr
}

// Everything in Lisp is an S-Expression. An SExp can
// be an atom or a list
type SExpr interface {
	isSExpr()
	Value() interface{}
}

////////////////////////////////////////////////////////////////////////////////
// Evaluate
////////////////////////////////////////////////////////////////////////////////
func eval(sexpr SExpr) SExpr {
	switch sexpr.(type) {
	case SExprList:
		fmt.Println("Found a list ", sexpr)
		list := sexpr.(SExprList)
		f := list.Expr[0]
		switch f.(type) {
		case Symbol:
			break
		default:
			panic("Only symbols are allowed as functions")
		}
		fn := list.Expr[0].(Symbol)
		switch fn.Name {
		case "+":
			fmt.Println("+ is a special form")
			add := func(args []SExpr) SExpr {
				fmt.Println("ARgs ", args)
				var val float32 = 0.0
				for _, arg := range args {
					argValue := arg.Value()
					switch argValue.(type) {
					case int:
						val += float32(argValue.(int))
						break
					case float32:
						val += float32(argValue.(float32))
						break
					default:
						var x bool = argValue.(bool)
						fmt.Println(x)
						panic("Value must be int or float32")
					}
				}
				return F32{val}
			}

			argValues := make([]SExpr, 0)
			for _, v := range list.Expr[1:] {
				argValues = append(argValues, eval(v.(SExpr)))
			}
			return add(argValues)
		default:
			panic("Unknown function")
		}
		break
	case Atom:
		fmt.Println("Atom")
		return sexpr
	default:
		panic("Unknown type")
	}
	return I32{0}
}

func main() {
	sexpr1 := NewList([]SExpr{S("+"), I32{3}, I32{2}})
	sexpr := NewList([]SExpr{S("+"), I32{3}, I32{2}, sexpr1})
	fmt.Println(sexpr)
	fmt.Println(eval(sexpr))
}
