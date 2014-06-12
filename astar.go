package main

import "container/heap"

func Manhattan(n1, n2 *Cell) int {
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
	closed := make(map[Cord]*Cell)

	G := make(map[Cord]int)
	H := make(map[Cord]int)
	G[start.Loc] = 0
	H[start.Loc] = Manhattan(start, goal)
	start.Key = H[start.Loc]
	parent := make(map[Cord]*Cell)

	heap.Init(open)
	heap.Push(open, start)

	for curr := heap.Pop(open).(*Cell); curr.Loc != goal.Loc; curr = heap.Pop(open).(*Cell) {
		closed[curr.Loc] = curr
		newCost := G[curr.Loc] + 1
		for _, adj := range curr.Adj {
			H[adj.Loc] = Manhattan(adj, goal)
			if !adj.Walkable() {
				continue
			}
			if _, inClosed := closed[adj.Loc]; inClosed {
				continue
			} else if inOpen := open.Contains(adj); !inOpen || newCost < G[adj.Loc] {
				parent[adj.Loc] = curr
				G[adj.Loc] = newCost
				adj.Key = G[adj.Loc] + H[adj.Loc]
				adj.Base = 'o'
				if inOpen {
					heap.Fix(open, adj.I)
				} else {
					heap.Push(open, adj)
				}
			}
		}
	}
	path := make([]*Cell, 0)
	for curr := parent[goal.Loc]; curr.Loc != start.Loc; curr = parent[curr.Loc] {
		path = append(path, curr)
	}
	return path
}
