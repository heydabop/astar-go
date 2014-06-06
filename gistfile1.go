package main

import (
	"container/heap"
	"fmt"
	"math/rand"
	"time"
)

type PQueue []*Node

type Cell struct {
	Top    *Cell
	Bottom *Cell
	Left   *Cell
	Right  *Cell
	Base   byte
	Unit   *Walker
	Row    int
	Col    int
}

type Grid struct {
	Grid [][]Cell
}

type Node struct {
	This   *Cell
	Parent *Cell
	F      int
	I      int
}

type Walker struct {
	Pos *Cell
	Val byte
}

func (pq PQueue) Len() int { return len(pq) }

func (pq PQueue) Less(i, j int) bool {
	return pq[i].F < pq[j].F
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

func (pq *PQueue) Update(node *Node, this, parent *Cell, f int) {
	heap.Remove(pq, node.I)
	node.This = this
	node.Parent = parent
	heap.Push(pq, node)
}

func (g Grid) Println() {
	for _, row := range g.Grid {
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

func (w *Walker) RandStep() {
	var oldCell = w.Pos
	var newCell *Cell
	for {
		switch rand.Intn(3) {
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
			continue
		} else {
			break
		}
	}
	oldCell.Unit = nil
	newCell.Unit = w
	w.Pos = newCell
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

func main() {
	rand.Seed(time.Now().Unix())
	rows, cols := 12, 24
	area := Grid{make([][]Cell, rows)}
	for i := 0; i < len(area.Grid); i++ {
		area.Grid[i] = make([]Cell, cols)
	}
	for i := 0; i < len(area.Grid); i++ {
		for j := 0; j < len(area.Grid[i]); j++ {
			area.Grid[i][j].Row = i
			area.Grid[i][j].Col = j
			if i == 0 {
				area.Grid[i][j].Top = nil
				area.Grid[i][j].Base = 'X'
			} else {
				area.Grid[i][j].Top = &area.Grid[i-1][j]
			}
			if i == len(area.Grid)-1 {
				area.Grid[i][j].Bottom = nil
				area.Grid[i][j].Base = 'X'
			} else {
				area.Grid[i][j].Bottom = &area.Grid[i+1][j]
			}
			if j == 0 {
				area.Grid[i][j].Left = nil
				area.Grid[i][j].Base = 'X'
			} else {
				area.Grid[i][j].Left = &area.Grid[i][j-1]
			}
			if j == len(area.Grid[i])-1 {
				area.Grid[i][j].Right = nil
				area.Grid[i][j].Base = 'X'
			} else {
				area.Grid[i][j].Right = &area.Grid[i][j+1]
			}
			if area.Grid[i][j].Base == 0 {
				area.Grid[i][j].Base = '.'
			}
		}
	}

	w := Walker{&area.Grid[rows/2][cols/2], 'O'}
	area.Grid[rows/2][cols/2].Unit = &w

	area.Grid[3][12].Base = 'X'

	area.Println()

}
