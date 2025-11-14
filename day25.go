package adventofcode2023

type Day25Puzzle map[string][]string

func NewDay25(lines []string) (Day25Puzzle, error) {
	puzzle := make(Day25Puzzle, 2000)

	for _, line := range lines {
		if line == "" {
			continue
		}

		// Parse format: "abc: def ghi jkl"
		colonIdx := -1
		for i := 0; i < len(line); i++ {
			if line[i] == ':' {
				colonIdx = i
				break
			}
		}
		if colonIdx == -1 {
			continue
		}

		from := line[:colonIdx]

		// Parse connections after ": "
		i := colonIdx + 2 // Skip ": "
		connections := make([]string, 0, 4)
		for i < len(line) {
			// Skip spaces
			for i < len(line) && line[i] == ' ' {
				i++
			}
			if i >= len(line) {
				break
			}

			// Find end of word
			start := i
			for i < len(line) && line[i] != ' ' {
				i++
			}
			connections = append(connections, line[start:i])
		}

		// Add bidirectional connections
		if puzzle[from] == nil {
			puzzle[from] = make([]string, 0, 8)
		}
		for _, to := range connections {
			puzzle[from] = append(puzzle[from], to)
			if puzzle[to] == nil {
				puzzle[to] = make([]string, 0, 8)
			}
			puzzle[to] = append(puzzle[to], from)
		}
	}

	return puzzle, nil
}

func Day25(puzzle Day25Puzzle, part1 bool) uint {
	// Find 3 edges to cut that split the graph into 2 components
	// Use edge betweenness: count edges in shortest paths between random pairs

	// Build edge usage map
	type edge struct{ a, b string }
	edgeCount := make(map[edge]int, 10000)

	// Get all nodes
	nodes := make([]string, 0, len(puzzle))
	for node := range puzzle {
		nodes = append(nodes, node)
	}

	// Sample shortest paths to find heavily used edges
	// Use efficient BFS with parent tracking
	for i := 0; i < len(nodes); i++ {
		parent := bfsParents(puzzle, nodes[i])

		// Reconstruct paths to all reachable nodes
		for j := i + 1; j < len(nodes); j++ {
			// Trace back from nodes[j] to nodes[i]
			curr := nodes[j]
			for parent[curr] != "" {
				prev := parent[curr]
				a, b := prev, curr
				if a > b {
					a, b = b, a
				}
				edgeCount[edge{a, b}]++
				curr = prev
			}
		}
	}

	// Find top 3 most used edges
	topEdges := make([]edge, 0, 3)
	for len(topEdges) < 3 {
		var maxEdge edge
		maxCount := 0
		for e, count := range edgeCount {
			if count > maxCount {
				maxEdge = e
				maxCount = count
			}
		}
		topEdges = append(topEdges, maxEdge)
		delete(edgeCount, maxEdge)
	}

	// Remove these 3 edges from graph
	graph := make(Day25Puzzle, len(puzzle))
	for node, neighbors := range puzzle {
		graph[node] = append([]string{}, neighbors...)
	}

	for _, e := range topEdges {
		removeEdge(graph, e.a, e.b)
	}

	// Count size of one component
	visited := make(map[string]bool, len(graph))
	var count int
	var dfs func(string)
	dfs = func(node string) {
		if visited[node] {
			return
		}
		visited[node] = true
		count++
		for _, neighbor := range graph[node] {
			dfs(neighbor)
		}
	}

	// Start DFS from any node
	for node := range graph {
		dfs(node)
		break
	}

	// Other component size
	other := len(graph) - count

	return uint(count * other)
}

func bfsParents(graph Day25Puzzle, start string) map[string]string {
	parent := make(map[string]string, len(graph))
	parent[start] = ""
	queue := make([]string, len(graph))
	queue[0] = start

	head, tail := 0, 1
	for head < tail {
		curr := queue[head]
		head++

		for _, neighbor := range graph[curr] {
			if _, seen := parent[neighbor]; !seen {
				parent[neighbor] = curr
				queue[tail] = neighbor
				tail++
			}
		}
	}

	return parent
}

func removeEdge(graph Day25Puzzle, a, b string) {
	graph[a] = filter(graph[a], b)
	graph[b] = filter(graph[b], a)
}

func filter(slice []string, remove string) []string {
	result := make([]string, 0, len(slice))
	for _, s := range slice {
		if s != remove {
			result = append(result, s)
		}
	}
	return result
}
