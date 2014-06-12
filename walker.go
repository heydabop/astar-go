package main

import (
	"fmt"
	"math/rand"
)

type Walker struct {
	Pos *Cell
	Val byte
}

func (w *Walker) Step(d int) bool {
	var oldCell = w.Pos
	var newCell *Cell
	switch d {
	case 0:
		newCell = w.Pos.Adj[0]
		fmt.Println("UP")
	case 1:
		newCell = w.Pos.Adj[1]
		fmt.Println("RIGHT")
	case 2:
		newCell = w.Pos.Adj[2]
		fmt.Println("DOWN")
	case 3:
		newCell = w.Pos.Adj[3]
		fmt.Println("LEFT")
	}
	if !newCell.Walkable() {
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
