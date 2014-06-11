package main

type PQueue []*Cell

func (pq PQueue) Len() int { return len(pq) }

func (pq PQueue) Less(i, j int) bool {
	return pq[i].Key < pq[j].Key
}

func (pq PQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].I = i
	pq[j].I = j
}

func (pq *PQueue) Push(x interface{}) {
	n := len(*pq)
	node := x.(*Cell)
	node.I = n
	(*pq) = append(*pq, node)
}

func (pq *PQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	node := old[n-1]
	node.I = -1
	*pq = old[:n-1]
	return node
}

func (pq PQueue) Contains(elem *Cell) bool {
	for _, node := range pq {
		if elem.Loc == node.Loc {
			return true
		}
	}
	return false
}
