type LinkNode struct {
  Next  *LinkNode
  Value *TreeNode
}

// Chained queue, first in first out
type LinkQueue struct {
  root *LinkNode  // Starting point of the chain table
  size int        // Number of elements in the queue
  lock sync.Mutex // Locks used for concurrent security
}

// in
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

//
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