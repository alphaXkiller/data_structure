package main

import "fmt"

type Node struct {
	data  int
	left  *Node
	right *Node
}

func createNode(data int) *Node {
	return &Node{
		data:  data,
		left:  nil,
		right: nil,
	}
}

func insert(root **Node, data int) {
	if *root == nil {
		*root = createNode(data)
		return
	}

	if data > (**root).data {
		// If it's the last node, create a new node, otherwise if looking for the last node
		if (**root).right == nil {
			(**root).right = createNode(data)
		} else {
			insert(&(**root).right, data)
		}
		return
	}

	// If it's the last node, create a new node, otherwise if looking for the last node
	if data < (**root).data {
		if (*root).left == nil {
			(*root).left = createNode(data)
		} else {
			insert(&(**root).left, data)
		}
		return
	}
}

func search(root *Node, data int) *Node {
	if root == nil {
		return nil
	}

	if root.data == data {
		return root
	}

	if data <= root.data {
		return search(root.left, data)
	}

	// data is larger than the root data
	return search(root.right, data)
}

func main() {
	var root *Node

	insert(&root, 15)
	insert(&root, 10)
	insert(&root, 20)
	insert(&root, 25)
	insert(&root, 8)
	insert(&root, 12)

	fmt.Println(root.data)
	fmt.Println(root.left.data)
	fmt.Println(root.right.data)
	fmt.Println(root.left.left.data)
	fmt.Println(root.left.right.data)
	fmt.Println(root.right.right.data)

	fmt.Println("===> search: ", search(root, 1))
}
