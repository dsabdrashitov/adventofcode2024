package graph

type Node[K comparable, N any, E any] interface {
	Compare(other N) int
	Key() K
	Edges() []E
}

type Edge[N any] interface {
	To() N
}
