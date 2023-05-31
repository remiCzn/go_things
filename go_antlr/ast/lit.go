package ast

import "fmt"

type Lit struct {
	Val int
}

func (l Lit) String() string {
	return fmt.Sprintf("Lit(%v)", l.Val)
}

func (l Lit) Value(environnement map[string]int) (int, map[string]int) {
	return l.Val, environnement
}
