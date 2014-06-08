package main

import (
	"fmt"
	"container/heap"
	"math/rand"
)

func (pq PQueue) Len() int { return len(pq) }

func (pq PQueue) Less(i, j int) bool {
	return pq[i].G+pq[i].H < pq[j].G+pq[j].H
}

func (pq PQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].I = i
	pq[j].I = j
}

func (pq *PQueue) Push(x interface{}) {
	n := len(*pq)
	node := x.(*Node)
	node.I = n
	*pq = append(*pq, node)
}

func (pq *PQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	node := old[n-1]
	node.I = -1
	*pq = old[:n-1]
	return node
}

func (pq *PQueue) Update(node *Node, this *Cell, parent *Node, f int) {
	heap.Remove(pq, node.I)
	node.This = this
	node.Parent = parent
	heap.Push(pq, node)
}

func (pq PQueue) Contains(elem *Node) bool {
	for _, node := range pq {
		if elem.This == node.This {
			return true
		}
	}
	return false
}

func (g Grid) Println() {
	for _, row := range g {
		for _, cell := range row {
			if cell.Unit == nil {
				fmt.Printf("%c", cell.Base)
			} else {
				fmt.Printf("%c", cell.Unit.Val)
			}
		}
		fmt.Println()
	}
}

func (w *Walker) Step(d int) bool {
	var oldCell = w.Pos
	var newCell *Cell
	switch d {
	case 0:
		newCell = w.Pos.Top
		fmt.Println("UP")
	case 1:
		newCell = w.Pos.Right
		fmt.Println("RIGHT")
	case 2:
		newCell = w.Pos.Bottom
		fmt.Println("DOWN")
	case 3:
		newCell = w.Pos.Left
		fmt.Println("LEFT")
	}
	if newCell.Base == 'X' {
		fmt.Println("BLOCKED")
		return false
	}
	oldCell.Unit = nil
	newCell.Unit = w
	w.Pos = newCell
	return true
}

func (w *Walker) RandStep() {
	for {
		walked := w.Step(rand.Intn(3))
		if walked {
			break
		}
	}
}

func (w *Walker) RandWalk(steps int) {
	for i := 0; i < steps; i++ {
		w.RandStep()
	}
}

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

//ive created a monster.

func CellsToNodes(g *Grid) map[Cell]Node {
	nodeMap := make(map[Cell]Node)
	for _, row := range *g {
		for _, cell := range row {
			nodeMap[cell] = Node{&cell,
				nil,
				[4]*Node{nil, nil, nil, nil},
				0, 0, 0}
		}
	}
	for _, node := range nodeMap {
		node.Adj[0] = nodeMap[*node.This.Top]
		node.Adj[1] = nodeMap[*node.This.Right]
		node.Adj[2] = nodeMap[*node.This.Bottom]
		node.Adj[3] = nodeMap[*node.This.Left]
	}
	return nodeMap
}
*/
