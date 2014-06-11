package main

import "container/heap"

func Manhattan(n1, n2 Cell) int {
	row := n2.Loc.Row - n1.Loc.Row
	col := n2.Loc.Col - n1.Loc.Col
	if row < 0 {
		row = -row
	}
	if col < 0 {
		col = -col
	}
	return row + col
}

func AStar(start, goal *Cell) []*Cell {
	open := &PQueue{}
	closed := make(map[Cord]Cell)

	G := make(map[Cord]int)
	H := make(map[Cord]int)
	parent := make(map[Cord]*Cell)

	heap.Init(open)
	heap.Push(open, start)

	for curr := heap.Pop(open).(*Cell); curr.Loc != goal.Loc; curr = heap.Pop(open).(*Cell) {
		closed[curr.Loc] = *curr
		cost := G[curr.Loc] + 1
		for _, adj := range curr.Adj {
			if open.Contains(adj) && cost < G[adj.Loc] {
				heap.Remove(open, adj.I)
			}
			if _, inClosed := closed[adj.Loc]; !open.Contains(adj) && !inClosed {
				G[adj.Loc] = cost
				H[adj.Loc] = Manhattan(*adj, *goal)
				adj.Key = G[adj.Loc] + H[adj.Loc]
				heap.Push(open, start)
				parent[adj.Loc] = curr
			}
		}
	}
	path := make([]*Cell, 12)
	for curr := parent[goal.Loc]; curr.Loc != start.Loc; curr = parent[curr.Loc] {
		path = append(path, curr)
	}
	return path
}
