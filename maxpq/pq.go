package maxpq

import (
	"github.com/kc596/UGCPriorityQueue/errors"
	"sync"
)

type PQ struct {
	n     int        // number of items on priority queue
	nodes []*Node    // store nodes at indices 0 to n-1
	lock  sync.Mutex // for accuracy in concurrency
}

/***************************************************************************
* Priority queue APIs
***************************************************************************/

// Returns a new max priority queue
func New() *PQ {
	return &PQ{
		n:    0,
		lock: sync.Mutex{},
	}
}

// Returns true if there are no nodes in priority queue
func (pq *PQ) IsEmpty() bool {
	pq.lock.Lock()
	defer pq.lock.Unlock()
	return pq.n == 0
}

// Returns the number of nodes in priority queue
func (pq *PQ) Size() int {
	pq.lock.Lock()
	defer pq.lock.Unlock()
	return pq.n
}

// Adds a new node to the priority queue
func (pq *PQ) Insert(x *Node) {
	pq.lock.Lock()
	defer pq.lock.Unlock()
	// double the size of array if necessary
	if pq.n >= len(pq.nodes) {
		pq.resize(2 * (len(pq.nodes) + 1))
	}
	pq.nodes[pq.n] = x
	pq.swim(pq.n)
	pq.n += 1
}

// Returns highest priority node of the priority queue
func (pq *PQ) Max() (*Node, error) {
	pq.lock.Lock()
	defer pq.lock.Unlock()
	if pq.n == 0 {
		return nil, errors.ErrNoSuchElement
	}
	return pq.nodes[0], nil
}

// Returns highest priority node and deletes it from priority queue
func (pq *PQ) Pop() (*Node, error) {
	pq.lock.Lock()
	defer pq.lock.Unlock()
	if pq.n == 0 {
		return nil, errors.ErrNoSuchElement
	}
	max := pq.nodes[0]
	pq.n -= 1
	pq.swap(0, pq.n)
	pq.nodes[pq.n] = nil
	pq.sink(0)
	// halve the size of array if necessary
	if (pq.n > 0) && (pq.n == len(pq.nodes)/4) {
		pq.resize(len(pq.nodes) / 2)
	}
	return max, nil
}

/***************************************************************************
* Helper functions to manage the heap
***************************************************************************/

func (pq *PQ) swim(x int) {
	for x > 0 {
		parent := (x - 1) >> 1
		if pq.nodes[parent].priority < pq.nodes[x].priority {
			pq.swap(x, parent)
			x = parent
		} else {
			break
		}
	}
}

func (pq *PQ) sink(x int) {
	for x < pq.n {
		child := x<<1 + 1
		if (child+1 < pq.n) && (pq.nodes[child+1].priority > pq.nodes[child].priority) {
			child += 1
		}
		if child >= pq.n {
			break
		}
		if pq.nodes[x].priority < pq.nodes[child].priority {
			pq.swap(x, child)
			x = child
		} else {
			break
		}
	}
}

/***************************************************************************
* Helper functions for swaps and resize
***************************************************************************/

func (pq *PQ) swap(x, y int) {
	temp := pq.nodes[x]
	pq.nodes[x] = pq.nodes[y]
	pq.nodes[y] = temp
}

func (pq *PQ) resize(capacity int) {
	temp := make([]*Node, capacity)
	for i := 0; i < pq.n; i++ {
		temp[i] = pq.nodes[i]
	}
	pq.nodes = temp
}
