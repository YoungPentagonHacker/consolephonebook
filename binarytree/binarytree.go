package binarytree

import (
	"fmt"

	db "github.com/YoungPentagonHacker/consolephonebook/database"
)

type Node struct {
	Value db.PersonDto
	Left  *Node
	Right *Node
}

func CreateTree(users []db.PersonDto) (root Node) {
	if len(users) != 0 {
		root = Node{users[0], nil, nil}
	}
	for i := 1; i < len(users); i++ {
		root.Add(users[i])
	}
	return root
}

// adds new node to the head
func (head *Node) Add(per db.PersonDto) {
	if head.Value.Name > per.Name {
		if head.Left == nil {
			head.Left = &Node{per, nil, nil}
		} else {
			head.Left.Add(per)
		}
		return
	}
	if head.Value.Name < per.Name {
		if head.Right == nil {
			head.Right = &Node{per, nil, nil}
		} else {
			head.Right.Add(per)
		}
		return

	}
	return
}

// prints tree by direct bypass
func (head Node) PrintTree() {
	fmt.Printf("Имя:%s  Номера:%s \n", head.Value.Name, head.Value.PhoneNumbers)
	if head.Left != nil {
		head.Left.PrintTree()
	}
	if head.Right != nil {
		head.Right.PrintTree()
	}
}

// returns GetDepth of the head
func (head Node) GetDepth() int {
	var res int
	var temp func(h Node, maxDepth *int, currentDepth int)
	temp = func(h Node, maxDepth *int, currentDepth int) {
		if h.Left != nil {
			currentDepth++
			temp(*h.Left, maxDepth, currentDepth)
		} else {
			if currentDepth > *maxDepth {
				*maxDepth = currentDepth
			}
			currentDepth = 1
		}
		if h.Right != nil {
			currentDepth++
			temp(*h.Right, maxDepth, currentDepth)
		} else {
			if currentDepth > *maxDepth {
				*maxDepth = currentDepth
			}
			currentDepth = 1
		}
	}
	temp(head, &res, 1)
	return res
}

// returns pointer of searched(by name) node  or nil if node hasn't been founded
func (head Node) FindByName(name string) *Node {
	if head.Value.Name == name {
		return &head
	}
	if head.Left != nil {
		if res := head.Left.FindByName(name); res != nil {
			return res
		}
	}
	if head.Right != nil {
		if res := head.Right.FindByName(name); res != nil {
			return res
		}
	}
	return nil
}

// returns deleted node or nil if there was not any node with this value
func (head *Node) DeleteNode(name string) *Node {

	if head.Left != nil {
		if head.Left.Value.Name == name {
			temp := head.Left
			head.Left = nil
			return temp
		}
		if res := head.Left.DeleteNode(name); res != nil {

			return res
		}
	}
	if head.Right != nil {
		if head.Right.Value.Name == name {
			temp := head.Right
			head.Right = nil
			return temp

		}
		if res := head.Right.DeleteNode(name); res != nil {
			return res
		}
	}
	return nil
}
