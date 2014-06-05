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
}
