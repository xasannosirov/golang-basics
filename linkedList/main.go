package main

import (
	"container/list"
	"fmt"
)

// add element
func Push(elem interface{}, queue *list.List) {
	queue.PushBack(elem)
}

// remove an item and return it
func Pop(queue *list.List) interface{} {
	res := queue.Remove(queue.Front())
	return res
}

// printing without spaces
func printQueue(queue *list.List) {
	for temp := queue.Front(); temp != nil; temp = temp.Next() {
		fmt.Printf("%v ", temp.Value)
	}
	fmt.Println()
}

// reverse the list
func ReverseList(queue *list.List) *list.List {
	reversedList := list.New()
	for elem := queue.Back(); elem != nil; elem = elem.Prev() {
		reversedList.PushBack(elem.Value)
	}
	return reversedList
}

func main(){
	// create a new linked list
	new := list.New()
	// add element
	Push(1, new)
	Push(2, new)
	printQueue(new)
	// remove element and print element
	rem := Pop(new)
	fmt.Println(rem)
	// reverse linked list
	copy := ReverseList(new)
	printQueue(copy)
}