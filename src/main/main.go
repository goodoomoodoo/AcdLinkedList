package main

import (
	"fmt"
	"../LinkedList"
)

func main() {
	TestList0 := LinkedList.NewLinkedList()
	fmt.Println("Hello")
	TestList0.Insert( 1 )
	fmt.Println( TestList0.Header.Element == nil )
	fmt.Println( *TestList0.Header.Next.Element )
}