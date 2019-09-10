package classes

import (
	"strconv"
)

type List struct {
	First *Node
}

func NewList(length int) List {
	list := List{}
	nodes := make([]Node, length)
	for i := 0; i < len(nodes); i++ {
		nodes[i].Data = i
	}
	for i := len(nodes) - 2; i >= 0; i-- {
		next := nodes[i+1]
		nodes[i].Next = &next
	}
	first := nodes[0]
	list.First = &first
	return list
}

func (list *List) ToString() string {
	result := ""
	node := list.First
	for true {
		result += " " + strconv.Itoa(node.Data)
		if node.Next != nil {
			node = node.Next
		} else {
			break
		}
	}
	return result
}

func (list *List) Reverse() {
	current := list.First
	var prev *Node
	prev = nil
	for true {
		next := current.Next
		current.Next = prev
		if next != nil {
			prev = current
			current = next
		} else {
			break
		}
	}
	list.First = current
}
