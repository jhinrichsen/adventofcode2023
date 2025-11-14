package adventofcode2023

type priorityItem interface {
	GetIndex() int
	SetIndex(int)
}

type state17 struct {
	row, col    int
	dr, dc      int
	consecutive int
	heat        uint
	index       int
}

func (s *state17) GetIndex() int    { return s.index }
func (s *state17) SetIndex(idx int) { s.index = idx }

type priorityQueue[T priorityItem] struct {
	items []T
	less  func(T, T) bool
}

func newPriorityQueue[T priorityItem](less func(T, T) bool) *priorityQueue[T] {
	return &priorityQueue[T]{items: make([]T, 0), less: less}
}

func (pq *priorityQueue[T]) Len() int { return len(pq.items) }
func (pq *priorityQueue[T]) Less(i, j int) bool {
	return pq.less(pq.items[i], pq.items[j])
}
func (pq *priorityQueue[T]) Swap(i, j int) {
	pq.items[i], pq.items[j] = pq.items[j], pq.items[i]
	pq.items[i].SetIndex(i)
	pq.items[j].SetIndex(j)
}
func (pq *priorityQueue[T]) Push(x interface{}) {
	n := len(pq.items)
	item := x.(T)
	item.SetIndex(n)
	pq.items = append(pq.items, item)
}
func (pq *priorityQueue[T]) Pop() interface{} {
	old := pq.items
	n := len(old)
	item := old[n-1]
	var zero T
	old[n-1] = zero
	item.SetIndex(-1)
	pq.items = old[0 : n-1]
	return item
}
