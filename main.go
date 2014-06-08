package main

import (
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	rows, cols := 12, 24
	area := make(Grid, rows)
	for i := 0; i < len(area); i++ {
		area[i] = make([]Cell, cols)
	}
	for i := 0; i < len(area); i++ {
		for j := 0; j < len(area[i]); j++ {
			area[i][j].Row = i
			area[i][j].Col = j
			if i == 0 {
				area[i][j].Top = nil
				area[i][j].Base = 'X'
			} else {
				area[i][j].Top = &area[i-1][j]
			}
			if i == len(area)-1 {
				area[i][j].Bottom = nil
				area[i][j].Base = 'X'
			} else {
				area[i][j].Bottom = &area[i+1][j]
			}
			if j == 0 {
				area[i][j].Left = nil
				area[i][j].Base = 'X'
			} else {
				area[i][j].Left = &area[i][j-1]
			}
			if j == len(area[i])-1 {
				area[i][j].Right = nil
				area[i][j].Base = 'X'
			} else {
				area[i][j].Right = &area[i][j+1]
			}
			if area[i][j].Base == 0 {
				area[i][j].Base = '.'
			}
		}
	}

	w := Walker{&area[rows/2][cols/2], 'O'}
	area[rows/2][cols/2].Unit = &w

	area[3][12].Base = 'X'

	for i := 0; i < 10; i++ {
		w.RandWalk(3)
		area.Println()
	}

}
