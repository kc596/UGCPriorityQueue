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

// Populate target with value of node - must be of original type
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
