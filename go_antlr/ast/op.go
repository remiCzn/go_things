package ast

import "fmt"

type Op int

const (
	PLUS    Op = 0
	MINUS   Op = 1
	TIMES   Op = 2
	DIVIDES Op = 3
)

type Bop struct {
	T1 Term
	Op Op
	T2 Term
}

func (t Bop) String() string {
	var op string
	switch t.Op {
	case PLUS:
		op = "+"
	case MINUS:
		op = "-"
	case TIMES:
		op = "x"
	case DIVIDES:
		op = "/"
	default:
		panic("Should not happen: unknown operator")
	}
	return fmt.Sprintf("BOp(%v,%v,%v)", t.T1, op, t.T2)
}

func (t Bop) Value(environnement map[string]int) (int, map[string]int) {
	t1, _ := t.T1.Value(environnement)
	t2, _ := t.T2.Value(environnement)
	var value int
	switch t.Op {
	case PLUS:
		value = t1 + t2
	case MINUS:
		value = t1 - t2
	case TIMES:
		value = t1 * t2
	case DIVIDES:
		value = t1 / t2
	default:
		panic("Should not happen: unknown operator")
	}
	return value, environnement
}
