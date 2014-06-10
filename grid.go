package main

import "fmt"

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
