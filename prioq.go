package adventofcode2023

import "cmp"

type priorityItem[P cmp.Ordered] interface {
	Priority() P
	GetIndex() int
	SetIndex(int)
}

type priorityQueue[P cmp.Ordered, T priorityItem[P]] struct {
	items []T
}

func newPriorityQueue[P cmp.Ordered, T priorityItem[P]]() *priorityQueue[P, T] {
	return &priorityQueue[P, T]{items: make([]T, 0)}
}

func (pq *priorityQueue[P, T]) Len() int { return len(pq.items) }
func (pq *priorityQueue[P, T]) Less(i, j int) bool {
	return pq.items[i].Priority() < pq.items[j].Priority()
}
func (pq *priorityQueue[P, T]) Swap(i, j int) {
	pq.items[i], pq.items[j] = pq.items[j], pq.items[i]
	pq.items[i].SetIndex(i)
	pq.items[j].SetIndex(j)
}
func (pq *priorityQueue[P, T]) Push(x interface{}) {
	n := len(pq.items)
	item := x.(T)
	item.SetIndex(n)
	pq.items = append(pq.items, item)
}
func (pq *priorityQueue[P, T]) Pop() interface{} {
	old := pq.items
	n := len(old)
	item := old[n-1]
	var zero T
	old[n-1] = zero
	item.SetIndex(-1)
	pq.items = old[0 : n-1]
	return item
}
