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
				if !cell.Walkable() {
					fmt.Printf("\x1b[31m")
				} else if cell.Base == '+' {
					fmt.Printf("\x1b[32m")
				} else if cell.Base == '*' || cell.Base == '-' {
					fmt.Printf("\x1b[34m")
				}
				fmt.Printf("%c", cell.Base)
				fmt.Printf("\x1b[0m")
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
