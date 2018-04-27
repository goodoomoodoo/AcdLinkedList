package LinkedList

import "fmt"

type LinkedList struct{
	Header *ListNode
	size int
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
	newList := LinkedList{ header , 0 }
	return &newList
}

/**
*	Insert things to the list
**/
func ( lst *LinkedList ) Insert( newElement int , curNode *ListNode ) {
	//declaration for new Node
	newNode := ListNode{ newElement , nil }

	newNode.Next = curNode.Next
	curNode.Next = &newNode

	lst.size++
}

/**
*	Add element to the list in ascending order
*	@param int element getting added to the list
**/
func ( lst *LinkedList ) Add( newElement int ) {
	nodePtr := lst.Header

	for ; nodePtr.Next != nil && nodePtr.Next.Element < newElement ;
		 nodePtr = nodePtr.Next { }

	lst.Insert( newElement , nodePtr )
}

/**
*	Remove element from the list
*	@param int element getting remove from the list
*	@return true if the element is successfully removed
**/
func ( lst *LinkedList ) Remove( oldElement int ) bool {
	if lst.IsEmpty( ){
		return false
	}

	nodePtr := lst.Header

	for ; nodePtr.Next.Element != oldElement ;
		nodePtr = nodePtr.Next {
			if nodePtr.Next.Next == nil{
				return false
			}
	}

	nodePtr.Next = nodePtr.Next.Next
	lst.size--

	return true
}

/**
*	Return the element at the given index
*	@param int index 
*	@return the node of the index if the node exist
**/
func ( lst *LinkedList ) Retrieve( index int ) *ListNode {
	if index >= lst.size {
		return nil 
	}

	nodePtr := lst.Header

	for ; index > 0 ; index-- {	
		nodePtr = nodePtr.Next
	}

	return nodePtr
}

/**
*	Check if the list contains some element
*	@param int some element to be found
*	@return true if the item does exist in the list
**/
func ( lst *LinkedList ) Contains( someElement int ) bool {
	if lst.IsEmpty( ) {
		return false
	}

	nodePtr := lst.Header

	for ; nodePtr.Next.Element != someElement ; 
		nodePtr = nodePtr.Next { 
			if nodePtr.Next.Next == nil {
				return false
			}
	}

	return true
}

/**
*	Lazily delete the list
**/
func ( lst *LinkedList ) Clear( ) {
	lst.Header.Next = nil
	lst.size = 0
}

/**
*	@return the size of the list
**/
func ( lst *LinkedList ) Size( ) int {
	return lst.size
}

/**
*	Check if the list is empty
*	@return true if the list contain no element
**/
func ( lst *LinkedList ) IsEmpty( ) bool {
	return lst.Header.Next == nil
}

/**
*	Print every item in the current list
**/
func ( lst *LinkedList ) Show() {
	for nodePtr := lst.Header.Next ; nodePtr != nil ; nodePtr = nodePtr.Next {
		fmt.Printf( "%v " , nodePtr.Element )
	}
}