package slices

type Stack[T any] []T

func (r *Stack[T]) Pop() T {
	l := len(*r)
	x := (*r)[l-1]
	*r = (*r)[:l-1]
	return x
}
