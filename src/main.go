package main

import (
	"fmt"
	"sync"
)

type TreeNode struct {
	Data  int       // Used to store data
	Left  *TreeNode // left subtree
	Right *TreeNode // Right subtree
}

type LinkNode struct {
	Next  *LinkNode
	Value *TreeNode
}

// LinkQueue Chained queue, first in first out
type LinkQueue struct {
	root *LinkNode  // Starting point of the chain table
	size int        // Number of elements in the queue
	lock sync.Mutex // Locks used for concurrent security
}

// Add in
func (queue *LinkQueue) Add(v *TreeNode) {
	queue.lock.Lock()
	defer queue.lock.Unlock()
	// If the top of the stack is empty, then add the node
	if queue.root == nil {
		queue.root = new(LinkNode)
		queue.root.Value = v
	} else {
		// Otherwise the new element is inserted at the end of the chain
		newNode := new(LinkNode)
		newNode.Value = v
		// Traverse all the way to the end of the chain
		nowNode := queue.root

		for nowNode.Next != nil {
			nowNode = nowNode.Next
		}

		// The new node is placed at the end of the chain
		nowNode.Next = newNode
	}

	// Number of elements in the queue +1
	queue.size = queue.size + 1
}

// Remove
func (queue *LinkQueue) Remove() *TreeNode {
	queue.lock.Lock()
	defer queue.lock.Unlock()

	// empty queue
	if queue.size == 0 {
		panic("over limit")
	}

	// pop the top element
	topNode := queue.root
	v := topNode.Value
	// Chain the top element's successor links
	queue.root = topNode.Next
	// Number of elements in the queue -1
	queue.size = queue.size - 1
	return v
}

func (root *TreeNode) GetTreeNodeNum() int {
	if root == nil {
		return 0
	} else {
		// calculate the number of nodes under the left node
		// calculate the number of nodes under the right node
		// and finally add the number of root nodes.
		return root.Left.GetTreeNodeNum() + root.Right.GetTreeNodeNum() + 1
	}
}

func CreateNode(v int) *TreeNode {
	return &TreeNode{v, nil, nil}
}

func (root *TreeNode) GetTreeDegree() int {
	maxDegree := 0

	if root == nil {
		return maxDegree
	}

	if root.Left.GetTreeDegree() > root.Right.GetTreeDegree() {
		maxDegree = root.Left.GetTreeDegree()
	} else {
		maxDegree = root.Right.GetTreeDegree()
	}

	return maxDegree + 1
}

func (root *TreeNode) PreOrder() {
	if root != nil {
		// print root
		fmt.Print(root.Data, " ")
		// print left tree
		root.Left.PreOrder()
		// print right tree
		root.Right.PreOrder()
	}
}

func (root *TreeNode) PostOrder() {
	if root != nil {
		// print left tree
		root.Left.PostOrder()
		// print right tree
		root.Right.PostOrder()
		// print root
		fmt.Print(root.Data, " ")
	}
}

func (root *TreeNode) MidOrder() {
	if root != nil {
		// print left tree
		root.Left.MidOrder()
		// print root
		fmt.Print(root.Data, " ")
		// print right tree
		root.Right.MidOrder()
	}
}

func (root *TreeNode) LayerOrder() {
	if root == nil {
		return
	}
	// new a queue
	queue := new(LinkQueue)
	// add root
	queue.Add(root)

	for queue.size > 0 {
		// Constantly out of queue
		element := queue.Remove()

		fmt.Print(element.Data, " ")

		// The left subtree is not empty and is in the queue
		if element.Left != nil {
			queue.Add(element.Left)
		}

		// The right subtree is not empty and is in the queue
		if element.Right != nil {
			queue.Add(element.Right)
		}
	}
}

func InitTree() *TreeNode {
	root := CreateNode(1)           //root node
	root.Left = CreateNode(2)       //left subtree
	root.Right = CreateNode(3)      //right subtree
	root.Left.Right = CreateNode(4) //right subtree of left subtree
	root.Right.Left = CreateNode(5) //left subtree of right subtree
	root.Left.Left = CreateNode(6)
	root.Right.Right = CreateNode(7)
	return root
}

func main() {
	root := InitTree()

	//a := root.GetTreeNodeNum()
	//b := root.GetTreeDegree()

	root.LayerOrder()
}
