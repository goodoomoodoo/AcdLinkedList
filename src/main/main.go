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
	fmt.Println( TestList0.Remove(10) )

	fmt.Println( TestList0.Header.Element == 0 )
	fmt.Println( TestList0.Header.Next.Element )
	fmt.Println( "Cotains 11 " , TestList0.Contains( 11 ) )
	fmt.Println( "Size: " , TestList0.Size( ) )
	fmt.Println( "List is Empty " , TestList0.IsEmpty( ) )
	TestList0.Show()
	fmt.Println( )
	TestList0.Clear()
	fmt.Println( "List is Empty " , TestList0.IsEmpty( ) )
}