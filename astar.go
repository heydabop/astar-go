package main

func Manhattan(n1, n2 Cell) int {
	row := n2.Row - n1.Row
	col := n2.Col - n1.Col
	if row < 0 {
		row = -row
	}
	if col < 0 {
		col = -col
	}
	return row + col
}

/*
func AStar(start, goal *Node) []*Node {
	open := &PQueue{}
	closed := make(map[*Node]*Node)
	heap.Init(open)
	heap.Push(open, start)
	for curr := heap.Pop(open).(*Node); curr.This != goal.This; curr = heap.Pop(open).(*Node) {
		closed[curr] = curr
		cost := curr.G + 1
		for _, adj := range curr.Adj {
			if open.Contains(adj) && cost < adj.G {
				heap.Remove(open, adj.I)
			}
			if _, inClosed := closed[adj]; !open.Contains(adj) && !inClosed {
				adj.G = cost
				heap.Push(open, start)
				adj.H = Manhattan(*adj.This, *goal.This)
				adj.Parent = curr
			}

		}
	}
	return make([]*Node, 4)
}
*/
