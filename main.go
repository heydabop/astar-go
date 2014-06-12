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
			area[i][j].Loc.Row = i
			area[i][j].Loc.Col = j
			if i == 0 {
				area[i][j].Adj[0] = nil
				area[i][j].Base = 'X'
			} else {
				area[i][j].Adj[0] = &area[i-1][j]
			}
			if i == len(area)-1 {
				area[i][j].Adj[2] = nil
				area[i][j].Base = 'X'
			} else {
				area[i][j].Adj[2] = &area[i+1][j]
			}
			if j == 0 {
				area[i][j].Adj[3] = nil
				area[i][j].Base = 'X'
			} else {
				area[i][j].Adj[3] = &area[i][j-1]
			}
			if j == len(area[i])-1 {
				area[i][j].Adj[1] = nil
				area[i][j].Base = 'X'
			} else {
				area[i][j].Adj[1] = &area[i][j+1]
			}
			if area[i][j].Base == 0 {
				area[i][j].Base = '.'
			}
		}
	}

	/*w := Walker{&area[rows/2][cols/2], 'O'}
	area[rows/2][cols/2].Unit = &w

	area[3][12].Base = 'X'

	for i := 0; i < 10; i++ {
		w.RandWalk(3)
		area.Println()
	}*/

	area[3][8].Base = 'X'
	area[4][8].Base = 'X'
	area[5][8].Base = 'X'
	area[6][8].Base = 'X'
	area[7][8].Base = 'X'
	area[8][8].Base = 'X'
	area[9][8].Base = 'X'

	path := AStar(&area[3][4], &area[7][15])
	area[3][4].Base = '*'
	for _, cell := range path {
		cell.Base = '+'
	}
	area[7][15].Base = '-'
	area.Println()

}
