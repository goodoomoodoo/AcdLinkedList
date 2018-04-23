package main

import (
	"fmt"
	"../LinkedList"
	"math/rand"
)

func main() {
	TestList0 := LinkedList.NewLinkedList()
	fmt.Println("Hello")
	
	for i := 0 ; i < 20 ; i++ {
		TestList0.Add( rand.Intn(100) )
	}

	fmt.Println( TestList0.Header.Element == 0 )
	fmt.Println( TestList0.Header.Next.Element )
	TestList0.Show()
}