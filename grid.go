package main

import "fmt"

type Cord struct {
	Row int
	Col int
}

type Cell struct {
	Adj      [4]*Cell
	Base     byte
	Unit     *Walker
	Loc      Cord
	I        int
	Key      int
}

type Grid [][]Cell

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

func (c Cell) Walkable() bool {
	return c.Base != 'X'
}
