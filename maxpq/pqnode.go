package maxpq

import (
	"reflect"
)

type Node struct {
	value    interface{}
	priority float64
}

func NewNode(value interface{}, priority float64) *Node {
	return &Node{value: value, priority: priority}
}

/***************************************************************************
* Getters
***************************************************************************/

func (node *Node) GetPriority() float64 {
	return node.priority
}

func (node *Node) GetStringValue() string {
	return node.value.(string)
}

func (node *Node) GetFloat64Value() float64 {
	return node.value.(float64)
}

func (node *Node) GetFloat32Value() float32 {
	return node.value.(float32)
}

func (node *Node) GetIntValue() int {
	return node.value.(int)
}

func (node *Node) GetInt64Value() int64 {
	return node.value.(int64)
}

func (node *Node) GetInt32Value() int32 {
	return node.value.(int32)
}

func (node *Node) GetUIntValue() uint {
	return node.value.(uint)
}

func (node *Node) GetUInt64Value() uint64 {
	return node.value.(uint64)
}

func (node *Node) GetUInt32Value() uint32 {
	return node.value.(uint32)
}

func (node *Node) GetFuncValue() func() {
	return node.value.(func())
}

// Populate target with value of node - must be of original type
// This method has cost of reflection associated with it
// Recommended to be used only when type of value is not primitive
func (node *Node) GetValue(target interface{}) {
	targetPtr := reflect.Indirect(reflect.ValueOf(target))
	targetType := targetPtr.Type()
	sourceType := reflect.TypeOf(node.value)
	sourceVal := reflect.ValueOf(node.value)

	for i := 0; i < sourceType.NumField(); i++ {
		fieldName := sourceType.Field(i).Name
		_, sourceFieldExistInTarget := targetType.FieldByName(fieldName)
		if sourceFieldExistInTarget {
			targetPtr.FieldByName(fieldName).Set(sourceVal.FieldByName(fieldName))
		}
	}
}
