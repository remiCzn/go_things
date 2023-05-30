package main

type Term interface {
	value(environnement map[string]int) (int, map[string]int)
}

type Lit struct {
	Val int
}

func (l Lit) value(environnement map[string]int) (int, map[string]int) {
	return l.Val, environnement
}

type Op int

const (
	PLUS    Op = 0
	MINUS      = 1
	TIMES      = 2
	DIVIDES    = 3
)

type Bop struct {
	T1 Term
	Op Op
	T2 Term
}

func (t Bop) value(environnement map[string]int) (int, map[string]int) {
	t1, _ := t.T1.value(environnement)
	t2, _ := t.T2.value(environnement)
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

type Ifz struct {
	condTerm Term
	thenTerm Term
	elseTerm Term
}

func (t Ifz) value(environnement map[string]int) (int, map[string]int) {
	cond, _ := t.condTerm.value(environnement)
	if cond == 0 {
		return t.thenTerm.value(environnement)
	} else {
		return t.elseTerm.value(environnement)
	}
}

type Var struct {
	name string
}

func (t Var) value(environnement map[string]int) (int, map[string]int) {
	return environnement[t.name], environnement
}

type Let struct {
	varname    string
	assignTerm Term
}

func (t Let) value(environnement map[string]int) (int, map[string]int) {
	environnement[t.varname], _ = t.assignTerm.value(environnement)
	return -1, environnement
}
