package main

type LinkedList struct {
	head *Node
}

type Node struct {
	Value int
	Next  *Node
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func (ll *LinkedList) Get(indexTarget int) int {
	index := 0
	for node := ll.head; node != nil; node = node.Next {
		if index == indexTarget {
			return node.Value
		}
		index++
	}
	return -1
}

func (ll *LinkedList) InsertHead(val int) {
	newNode := &Node{Value: val}
	newNode.Next = ll.head
	ll.head = newNode
}

func (ll *LinkedList) InsertTail(val int) {
	for node := ll.head; node != nil; node = node.Next {
		if node.Next == nil {
			node.Next = &Node{Value: val}
			return
		}
	}
	ll.head = &Node{Value: val}
}

func (ll *LinkedList) Remove(indexTarget int) bool {
	if ll.head == nil {
		return false
	}
	if indexTarget == 0 {
		ll.head = ll.head.Next
		return true
	}
	index := 0
	prevNode := ll.head
	for node := ll.head; node != nil; node = node.Next {
		if index == indexTarget {
			prevNode.Next = node.Next
			return true
		}
		prevNode = node
		index++
	}
	return false
}

func (ll *LinkedList) GetValues() (ret []int) {
	for node := ll.head; node != nil; node = node.Next {
		ret = append(ret, node.Value)
	}
	return ret
}

func main() {

}
