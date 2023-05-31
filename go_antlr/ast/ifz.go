package ast

import "fmt"

type Ifz struct {
	CondTerm Term
	ThenTerm Term
	ElseTerm Term
}

func (t Ifz) String() string {
	return fmt.Sprintf("Ifz(%v,%v,%v)", t.CondTerm, t.ThenTerm, t.ElseTerm)
}

func (t Ifz) Value(environnement map[string]int) (int, map[string]int) {
	cond, _ := t.CondTerm.Value(environnement)
	if cond == 0 {
		return t.ThenTerm.Value(environnement)
	} else {
		return t.ElseTerm.Value(environnement)
	}
}
