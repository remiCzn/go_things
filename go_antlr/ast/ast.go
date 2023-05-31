package ast

type Term interface {
	Value(environnement map[string]int) (int, map[string]int)
}
