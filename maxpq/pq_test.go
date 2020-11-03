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

	node = NewNode("stringvalue", testPriority+1)
	assert.Equal("stringvalue", node.GetStringValue())
	assert.Equal(testPriority+1, node.GetPriority())

	float64value := rand.Float64()
	node = NewNode(float64value, testPriority+2)
	assert.Equal(float64value, node.GetFloat64Value())
	assert.Equal(testPriority+2, node.GetPriority())

	float32value := rand.Float32()
	node = NewNode(float32value, testPriority+3)
	assert.Equal(float32value, node.GetFloat32Value())
	assert.Equal(testPriority+3, node.GetPriority())

	intvalue := rand.Int()
	node = NewNode(intvalue, testPriority+4)
	assert.Equal(intvalue, node.GetIntValue())
	assert.Equal(testPriority+4, node.GetPriority())

	int64value := rand.Int63()
	node = NewNode(int64value, testPriority+5)
	assert.Equal(int64value, node.GetInt64Value())
	assert.Equal(testPriority+5, node.GetPriority())

	int32value := rand.Int31()
	node = NewNode(int32value, testPriority+6)
	assert.Equal(int32value, node.GetInt32Value())
	assert.Equal(testPriority+6, node.GetPriority())

	uintvalue := uint(rand.Uint32())
	node = NewNode(uintvalue, testPriority+7)
	assert.Equal(uintvalue, node.GetUIntValue())
	assert.Equal(testPriority+7, node.GetPriority())

	uint64value := rand.Uint64()
	node = NewNode(uint64value, testPriority+8)
	assert.Equal(uint64value, node.GetUInt64Value())
	assert.Equal(testPriority+8, node.GetPriority())

	uint32value := rand.Uint32()
	node = NewNode(uint32value, testPriority+9)
	assert.Equal(uint32value, node.GetUInt32Value())
	assert.Equal(testPriority+9, node.GetPriority())

	x := 1
	funcValue := func() { x++ }
	node = NewNode(funcValue, testPriority+10)
	fn := node.GetFuncValue()
	fn()
	assert.Equal(2, x)
	assert.Equal(testPriority+10, node.GetPriority())
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

func TestPQ_Clear(t *testing.T) {
	assert := assert.New(t)
	var nodes []*Node
	for i := 0; i < 100; i++ {
		nodeVal := nodeValue{K1: rand.Int(), K2: "Index" + fmt.Sprint(i), K3: struct{ K4 int }{K4: rand.Int()}}
		nodes = append(nodes, NewNode(nodeVal, float64(rand.Float64())))
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
	pq.Clear()
	assert.Zero(pq.Size())
	assert.True(pq.IsEmpty())
	assert.Equal(0, pq.n)

	// reinsert
	for _, node := range nodes {
		go pq.Insert(node)
	}
	for pq.Size() < len(nodes) {
		time.Sleep(10 * time.Millisecond)
	}
	newMax, err := pq.Pop()
	assert.Nil(err)
	assert.Equal(max, newMax)
	assert.Equal(len(nodes)-1, pq.Size())
	assert.False(pq.IsEmpty())
}

func TestPQ_ClearEmptyPQ(t *testing.T) {
	assert := assert.New(t)
	pq := New()
	pq.Clear()
	var nodes []*Node
	for i := 0; i < 100; i++ {
		priority := 10*i - i*i // 0, 9, 16, 21, 24, 25, 24, 21, ... Max value at 4
		nodeVal := nodeValue{K1: priority, K2: "Index" + fmt.Sprint(i), K3: struct{ K4 int }{K4: rand.Int()}}
		nodes = append(nodes, NewNode(nodeVal, float64(priority)))
	}
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
