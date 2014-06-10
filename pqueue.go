package main

import "container/heap"

type PQueue []*Node

func (pq PQueue) Len() int { return len(pq) }

func (pq PQueue) Less(i, j int) bool {
	return pq[i].G+pq[i].H < pq[j].G+pq[j].H
}

func (pq PQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].I = i
	pq[j].I = j
}

func (pq *PQueue) Push(x interface{}) {
	n := len(*pq)
	node := x.(*Node)
	node.I = n
	*pq = append(*pq, node)
}

func (pq *PQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	node := old[n-1]
	node.I = -1
	*pq = old[:n-1]
	return node
}

func (pq *PQueue) Update(node *Node, this *Cell, parent *Node, f int) {
	heap.Remove(pq, node.I)
	node.This = this
	node.Parent = parent
	heap.Push(pq, node)
}

func (pq PQueue) Contains(elem *Node) bool {
	for _, node := range pq {
		if elem.This == node.This {
			return true
		}
	}
	return false
}
