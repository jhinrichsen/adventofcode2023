package adventofcode2023

import "container/heap"

type state17 struct {
	row, col    int
	dr, dc      int
	consecutive int
	heat        uint
	index       int
}

type priorityQueue17 []*state17

func (pq priorityQueue17) Len() int { return len(pq) }
func (pq priorityQueue17) Less(i, j int) bool {
	return pq[i].heat < pq[j].heat
}
func (pq priorityQueue17) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}
func (pq *priorityQueue17) Push(x interface{}) {
	n := len(*pq)
	item := x.(*state17)
	item.index = n
	*pq = append(*pq, item)
}
func (pq *priorityQueue17) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

type visitKey17 struct {
	row, col    int
	dr, dc      int
	consecutive int
}

func Day17(lines []string, part1 bool) uint {
	if len(lines) == 0 {
		return 0
	}

	rows := len(lines)
	cols := len(lines[0])

	minConsecutive := 0
	maxConsecutive := 3
	if !part1 {
		minConsecutive = 4
		maxConsecutive = 10
	}

	visited := make(map[visitKey17]bool)
	pq := &priorityQueue17{}
	heap.Init(pq)

	heap.Push(pq, &state17{row: 0, col: 0, dr: 0, dc: 1, consecutive: 0, heat: 0})
	heap.Push(pq, &state17{row: 0, col: 0, dr: 1, dc: 0, consecutive: 0, heat: 0})

	for pq.Len() > 0 {
		current := heap.Pop(pq).(*state17)

		if current.row == rows-1 && current.col == cols-1 {
			if part1 || current.consecutive >= minConsecutive {
				return current.heat
			}
		}

		key := visitKey17{current.row, current.col, current.dr, current.dc, current.consecutive}
		if visited[key] {
			continue
		}
		visited[key] = true

		directions := [][2]int{}

		if current.consecutive < maxConsecutive && (current.dr != 0 || current.dc != 0) {
			directions = append(directions, [2]int{current.dr, current.dc})
		}

		if current.consecutive >= minConsecutive || (current.dr == 0 && current.dc == 0) {
			if current.dr == 0 {
				directions = append(directions, [2]int{-1, 0}, [2]int{1, 0})
			} else if current.dc == 0 {
				directions = append(directions, [2]int{0, -1}, [2]int{0, 1})
			} else {
				if current.dr != 0 {
					directions = append(directions, [2]int{0, -1}, [2]int{0, 1})
				} else {
					directions = append(directions, [2]int{-1, 0}, [2]int{1, 0})
				}
			}
		}

		for _, dir := range directions {
			newRow := current.row + dir[0]
			newCol := current.col + dir[1]

			if newRow < 0 || newRow >= rows || newCol < 0 || newCol >= cols {
				continue
			}

			newConsecutive := 1
			if dir[0] == current.dr && dir[1] == current.dc {
				newConsecutive = current.consecutive + 1
			}

			newHeat := current.heat + uint(lines[newRow][newCol]-'0')

			heap.Push(pq, &state17{
				row:         newRow,
				col:         newCol,
				dr:          dir[0],
				dc:          dir[1],
				consecutive: newConsecutive,
				heat:        newHeat,
			})
		}
	}

	return 0
}
