package ast

import "fmt"

type Var struct {
	Name string
}

func (t Var) String() string {
	return fmt.Sprintf("Var(%v)", t.Name)
}

func (t Var) Value(environnement map[string]int) (int, map[string]int) {
	return environnement[t.Name], environnement
}
