package main

import "container/heap"

type PQueue struct {
	val []*Node
	key map[*Node]int
}

func (pq PQueue) Len() int { return len(pq.val) }

func (pq PQueue) Less(i, j int) bool {
	return pq.key[pq.val[i]] < pq.key[pq.val[j]]
}

func (pq PQueue) Swap(i, j int) {
	pq.val[i], pq.val[j] = pq.val[j], pq.val[i]
	pq.val[i].I = i
	pq.val[j].I = j
}

func (pq *PQueue) Push(x interface{}) {
	n := len((*pq).val)
	node := x.(*Node)
	node.I = n
	(*pq).val = append((*pq).val, node)
}

func (pq *PQueue) Pop() interface{} {
	old := *pq
	n := len(old.val)
	node := old.val[n-1]
	node.I = -1
	(*pq).val = old.val[:n-1]
	return node
}

func (pq *PQueue) Update(node *Node, this *Cell, parent *Node, f int) {
	heap.Remove(pq, node.I)
	node.This = this
	node.Parent = parent
	heap.Push(pq, node)
}

func (pq PQueue) Contains(elem *Node) bool {
	for _, node := range pq.val {
		if elem.This == node.This {
			return true
		}
	}
	return false
}
