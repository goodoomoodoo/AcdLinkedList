package LinkedList

import "fmt"

type LinkedList struct{
	Header *ListNode
}

/**
*	List Node Type
**/
type ListNode struct {
	Element int
	Next *ListNode
}


/**
*	Constructor of LinkedList: not really necessary I guess
**/
func NewLinkedList() *LinkedList {
	//declare local variable header
	header := &ListNode{ 0 , nil }
	newList := LinkedList{ header }
	return &newList
}

/**
*	Insert things to the list
**/
func ( lst LinkedList ) Insert( newElement int , curNode *ListNode ) {
	//declaration for new Node
	newNode := ListNode{ newElement , nil }

	newNode.Next = curNode.Next
	curNode.Next = &newNode
}

/**
*	Add element to the list in ascending order
**/
func ( lst LinkedList ) Add( newElement int ) {
	nodePtr := lst.Header

	for ; nodePtr.Next != nil && nodePtr.Next.Element < newElement ;
		 nodePtr = nodePtr.Next {

	}

	lst.Insert( newElement , nodePtr )
}

/**
*	Print every item in the current list
**/
func ( lst LinkedList ) Show() {
	for nodePtr := lst.Header.Next ; nodePtr != nil ; nodePtr = nodePtr.Next {
		fmt.Printf( "%v " , nodePtr.Element )
	}
}