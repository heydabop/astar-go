package main

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

type Grid [][]Cell

type Node struct {
	This   *Cell
	Parent *Node
	Adj    [4]*Node
	G      int
	H      int
	I      int
}

type Walker struct {
	Pos *Cell
	Val byte
}
