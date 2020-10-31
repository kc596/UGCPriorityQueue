package maxpq

import (
	"fmt"
	"github.com/kc596/UGCPriorityQueue/errors"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
	"time"
)

type nodeValue struct {
	K1 int
	K2 string
	K3 struct {
		K4 int
	}
}

var (
	testNodeValue = nodeValue{K1: rand.Int(), K2: "test", K3: struct{ K4 int }{K4: rand.Int()}}
	testPriority  = float64(rand.Int())
	testNode      = NewNode(testNodeValue, testPriority)
)

func TestPQNode(t *testing.T) {
	assert := assert.New(t)
	node := testNode

	var nodeValueOut nodeValue
	node.GetValue(&nodeValueOut)
	assert.Equal(testNodeValue, nodeValueOut)
}

func TestPQ_Insert(t *testing.T) {
	assert := assert.New(t)
	pq := New()
	assert.Zero(pq.n)
	assert.Empty(pq.nodes)
	pq.Insert(testNode)
	assert.Equal(1, pq.n)
	assert.Greater(len(pq.nodes), 0)
	assert.Equal(testNode, pq.nodes[0])
}

func TestPQ_IsEmpty(t *testing.T) {
	assert := assert.New(t)
	pq := New()
	assert.True(pq.IsEmpty())
	pq.Insert(testNode)
	assert.False(pq.IsEmpty())
	pq.Pop()
	assert.True(pq.IsEmpty())
}

func TestPQ_Size(t *testing.T) {
	assert := assert.New(t)
	pq := New()
	assert.Zero(pq.Size())
	pq.Insert(testNode)
	assert.Equal(1, pq.Size())
	pq.Insert(NewNode(testNodeValue, testPriority+10))
	assert.Equal(2, pq.Size())
	pq.Pop()
	assert.Equal(1, pq.Size())
	pq.Pop()
	assert.Zero(pq.Size())
}

func TestPQ_Max(t *testing.T) {
	assert := assert.New(t)
	pq := New()
	pq.Insert(testNode)
	max, err := pq.Max()
	assert.Nil(err)
	assert.Equal(testNode, max)
	pq.Pop()
	max, err = pq.Max()
	assert.Nil(max)
	assert.EqualError(err, errors.ErrNoSuchElement.Error())
}

func TestPQ_Pop(t *testing.T) {
	assert := assert.New(t)
	pq := New()
	pq.Insert(testNode)
	max, err := pq.Pop()
	assert.Nil(err)
	assert.Equal(testNode, max)
	max, err = pq.Pop()
	assert.Nil(max)
	assert.EqualError(err, errors.ErrNoSuchElement.Error())
}

func TestPQConcurrency(t *testing.T) {
	assert := assert.New(t)
	var nodes []*Node
	for i := 0; i < 100; i++ {
		priority := 10*i - i*i // 0, 9, 16, 21, 24, 25, 24, 21, ... Max value at 4
		nodeVal := nodeValue{K1: priority, K2: "Index" + fmt.Sprint(i), K3: struct{ K4 int }{K4: rand.Int()}}
		nodes = append(nodes, NewNode(nodeVal, float64(priority)))
	}
	pq := New()
	for _, node := range nodes {
		go pq.Insert(node)
	}
	for pq.Size() < len(nodes) {
		time.Sleep(10 * time.Millisecond)
	}
	max, err := pq.Max()
	assert.Nil(err)
	assert.Equal(nodes[5], max)
	for range nodes {
		go pq.Pop()
	}
	for pq.Size() > 0 {
		time.Sleep(10 * time.Millisecond)
	}
	assert.Zero(pq.Size())
}
