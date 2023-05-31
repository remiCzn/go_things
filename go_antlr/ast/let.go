package ast

import "fmt"

type Let struct {
	VarName    string
	AssignTerm Term
}

func (t Let) String() string {
	return fmt.Sprintf("Let(%v,%v)", t.VarName, t.AssignTerm)
}

func (t Let) Value(environnement map[string]int) (int, map[string]int) {
	environnement[t.VarName], _ = t.AssignTerm.Value(environnement)
	return -1, environnement
}
