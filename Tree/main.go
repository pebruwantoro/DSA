package main

import "fmt"

type Node struct {
	data int
	prev *Node
	next *Node
}

/*
	      T1
		/	\
       T2   T3
      / \   / \
	 T4 T5 T6 T7
*/

func (n *Node) Push(data int) {
	if data < n.data {
		// If the new data is less than the current node's data, go previous

		// If there is no previous node, create a new one
		if n.prev == nil {
			n.prev = &Node{data: data}
		} else {
			// If there is a previous node, recursively call Push on it
			n.prev.Push(data)
		}
	} else if data > n.data {
		// If the new data is greater than the current node's data, go next

		// If there is no next node, create a new one
		if n.next == nil {
			n.next = &Node{data: data}
		} else {
			// If there is a next node, recursively call Push on it
			n.next.Push(data)
		}
	}
	// If the new data is equal to the current node's data, do nothing
}

func (n *Node) Pop(data int) *Node {
	if n == nil {
		return nil
	}

	if data < n.data {
		n.prev = n.prev.Pop(data)
	} else if data > n.data {
		n.next = n.next.Pop(data)
	} else {

		if n.prev == nil {
			return n.next
		}

		if n.next == nil {
			return n.prev
		}

		temp := n.next
		for temp.prev != nil {
			temp = temp.prev
		}

		n.data = temp.data

		n.next = n.next.Pop(temp.data)
	}

	return n
}

func (n *Node) Print() {
	if n == nil {
		return
	}

	fmt.Println("Tree structure: ")
	printNode(n, "", true)
}

func printNode(node *Node, prefix string, isTail bool) {
	if node == nil {
		return
	}

	fmt.Print(prefix)
	if isTail {
		fmt.Print("└── ")
	} else {
		fmt.Print("├── ")
	}
	fmt.Println(node.data)

	childPrefix := prefix
	if isTail {
		childPrefix += "    "
	} else {
		childPrefix += "│   "
	}

	if node.prev != nil {
		isPrevTail := (node.next == nil)
		printNode(node.prev, childPrefix, isPrevTail)
	}
	if node.next != nil {
		printNode(node.next, childPrefix, true)
	}
}

func main() {
	root := &Node{data: 10}

	fmt.Println("====START====")

	fmt.Println("====PUSH====")
	root.Push(5)
	root.Push(15)
	root.Push(3)
	root.Push(7)
	root.Push(12)
	root.Push(18)
	fmt.Println("====END OF PUSH====")

	root.Print()
	// The tree structure is now:
	//         10
	//       /    \
	//      5     15
	//     / \   /  \
	//    3   7 12   18

	root = root.Pop(10)
	fmt.Println("====AFTER POP====")
	root.Print()

	// The tree structure is now:
	//         12
	//       /    \
	//      5     15
	//     / \      \
	//    3   7      18

	root = root.Pop(5)
	fmt.Println("====AFTER POP====")
	root.Print()

	// The tree structure is now:
	//         12
	//       /    \
	//      7     15
	//     /       \
	//    3         18

	root.Push(14)
	fmt.Println("====AFTER PUSH====")
	root.Print()

	// The tree structure is now:
	//         12
	//       /    \
	//      7     15
	//     /     /  \
	//    3     14   18
}
