## UGC Priority Queue

**U**nbounded **G**eneric **C**oncurrent Priority Queue in GoLang.

### Installation

> go get github.com/kc596/UGCPriorityQueue

### Quickstart

```go
import (
	"github.com/kc596/UGCPriorityQueue/maxpq"
	"fmt"
)

func ExamplePQ() {
	// create new max priority queue
	pq := maxpq.NewMaxPQ()
    
	// creating new pq node with priority 1
	node1 := maxpq.NewNode("Value1", 1)
    
	// value of node could be of any type, int here
	node2 := maxpq.NewNode(101, 10)
    
	// add the nodes to pq
	pq.Insert(node1)
	pq.Insert(node2)
    
	// get the node with highest priority
	highestPriorityNode, _ := pq.Max()
	fmt.Printf("%+v\n", highestPriorityNode)
    
	// pop the nodes with highest priority
	highestPriorityNode, _ = pq.Pop()
	nextHighest, _ := pq.Pop()
	fmt.Printf("%+v\n", highestPriorityNode) // value:101, priority:10
	fmt.Printf("%+v\n", nextHighest)         // value:Value1, priority:1
}
```

