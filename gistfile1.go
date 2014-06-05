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
	area := Grid{make([][]Cell, 12)}
	for i := 0; i < len(area.Grid); i++ {
		area.Grid[i] = make([]Cell, 12)
	}
	for i := 0; i < len(area.Grid); i++ {
		for j := 0; j < len(area.Grid[i]); j++ {
			if i == 0 {
				area.Grid[i][j].Top = nil
				area.Grid[i][j].Val = "X"
			} else {
				area.Grid[i][j].Top = &area.Grid[i-1][j]
			}
			if i == len(area.Grid)-1 {
				area.Grid[i][j].Bottom = nil
				area.Grid[i][j].Val = "X"
			} else {
				area.Grid[i][j].Bottom = &area.Grid[i+1][j]
			}
			if j == 0 {
				area.Grid[i][j].Left = nil
				area.Grid[i][j].Val = "X"
			} else {
				area.Grid[i][j].Left = &area.Grid[i][j-1]
			}
			if j == len(area.Grid[i])-1 {
				area.Grid[i][j].Right = nil
				area.Grid[i][j].Val = "X"
			} else {
				area.Grid[i][j].Right = &area.Grid[i][j+1]
			}
			if len(area.Grid[i][j].Val) == 0 {
				area.Grid[i][j].Val = "."
			}
		}
	}
	for _, row := range area.Grid {
		for _, cell := range row {
			fmt.Print(cell.Val)
		}
		fmt.Println()
	}
}
