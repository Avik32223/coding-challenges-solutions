package heap

type Node interface {
	Compare(Node) int
}
