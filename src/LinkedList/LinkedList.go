package LinkedList

type LinkedList struct{
	Header *ListNode
}

/**
*	List Node Type
**/
type ListNode struct {
	Element *interface{}
	Next *ListNode
}


/**
*	Constructor of LinkedList: not really necessary I guess
**/
func NewLinkedList() *LinkedList {
	header := &ListNode{ nil , nil }
	newList := LinkedList{ header }
	return &newList
}

/**
*	Add things to the list
**/
func ( lst LinkedList ) Insert( newElement interface{} ) {
	//declaration for new Node
	newNode := ListNode{ &newElement , nil }

	if lst.Header.Next == nil {
		lst.Header.Next = &newNode
	} else {
		newNode.Next = lst.Header.Next
		lst.Header.Next = &newNode
	}
}