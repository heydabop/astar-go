package main

import (
	"fmt"
)

type Cell struct {
	Top    *Cell
	Bottom *Cell
	Left   *Cell
	Right  *Cell
	Val    string
}

type Grid struct {
	Grid [][]Cell
}

func main() {
	fmt.Println("Test!\n")
	area := Grid{make([][]Cell, 12)}
	for i := 0; i < len(area.Grid); i++ {
		area.Grid[i] = make([]Cell, 12)
	}
	for i := 0; i < len(area.Grid); i++ {
		for j := 0; j < len(area.Grid[i]); j++ {
			area.Grid[i][j].Top = &area.Grid[i-1][j]
			area.Grid[i][j].Bottom = &area.Grid[i+1][j]
			area.Grid[i][j].Left = &area.Grid[i][j-1]
			area.Grid[i][j].Right = &area.Grid[i][j+1]
			area.Grid[i][j].Val = "."
		}
	}
}
