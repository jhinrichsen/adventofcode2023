package adventofcode2023

import "cmp"

type priorityQueue[P cmp.Ordered, T any] struct {
	items []T
	less  func(T, T) bool
}

func newPriorityQueue[P cmp.Ordered, T any](less func(T, T) bool) *priorityQueue[P, T] {
	return &priorityQueue[P, T]{items: make([]T, 0), less: less}
}

func (pq *priorityQueue[P, T]) Len() int {
	return len(pq.items)
}

func (pq *priorityQueue[P, T]) Push(item T) {
	pq.items = append(pq.items, item)
	pq.up(len(pq.items) - 1)
}

func (pq *priorityQueue[P, T]) Pop() T {
	n := len(pq.items) - 1
	pq.items[0], pq.items[n] = pq.items[n], pq.items[0]
	pq.down(0, n)
	item := pq.items[n]
	pq.items = pq.items[:n]
	return item
}

func (pq *priorityQueue[P, T]) up(j int) {
	for {
		i := (j - 1) / 2
		if i == j || !pq.less(pq.items[j], pq.items[i]) {
			break
		}
		pq.items[i], pq.items[j] = pq.items[j], pq.items[i]
		j = i
	}
}

func (pq *priorityQueue[P, T]) down(i, n int) {
	for {
		j1 := 2*i + 1
		if j1 >= n || j1 < 0 {
			break
		}
		j := j1
		if j2 := j1 + 1; j2 < n && pq.less(pq.items[j2], pq.items[j1]) {
			j = j2
		}
		if !pq.less(pq.items[j], pq.items[i]) {
			break
		}
		pq.items[i], pq.items[j] = pq.items[j], pq.items[i]
		i = j
	}
}
