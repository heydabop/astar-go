package main

type Node struct {
	This   *Cell
	Parent *Node
	Adj    [4]*Node
	G      int
	H      int
	I      int
}
/*
//ive created a monster.

func CellsToNodes(g *Grid) map[Cell]Node {
	nodeMap := make(map[Cell]Node)
	for _, row := range *g {
		for _, cell := range row {
			nodeMap[cell] = Node{&cell,
				nil,
				[4]*Node{nil, nil, nil, nil},
				0, 0, 0}
		}
	}
	for _, node := range nodeMap {
		node.Adj[0] = nodeMap[*node.This.Top]
		node.Adj[1] = nodeMap[*node.This.Right]
		node.Adj[2] = nodeMap[*node.This.Bottom]
		node.Adj[3] = nodeMap[*node.This.Left]
	}
	return nodeMap
}
*/
