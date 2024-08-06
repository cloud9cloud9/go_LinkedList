package linked

type LinkedInterface[T any] interface {
	AddFirst(el T)
	AddLast(el T)
	Get(el T) (index int)
	Remove(el T)
	PrintList()
}
