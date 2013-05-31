package main

import "fmt"

type Node struct {
    Next *Node
    Data interface{}
}

type LinkedList struct {
    Root *Node
}

func (list *LinkedList) Insert(data interface{}) {
    // Create a new node
    node := new(Node)
    node.Data = data

    if list.Root == nil {
        list.Root = node
        return
    }

    current := list.Root
    for current.Next != nil {
        current = current.Next
    }
    current.Next = node
}

func (list *LinkedList) Contains(data interface{}) bool {
    current := list.Root
    for current != nil {
        if current.Data == data {
            return true
        }
        current = current.Next
    }

    return false
}

func (list *LinkedList) Print() {
    current := list.Root
    for current != nil {
        fmt.Println(current.Data)
        current = current.Next
    }
}

func main() {
    list := new(LinkedList)
    list.Insert(1)
    list.Insert(2)
    list.Insert(3)
    list.Print()

    fmt.Println("Does 2 exist in the list?")
    if list.Contains(2) {
        fmt.Println("Yes!")
    } else {
        fmt.Println("No.")
    }
    fmt.Println("Does 7 exist in the list?")
    if list.Contains(7) {
        fmt.Println("Yes!")
    } else {
        fmt.Println("No.")
    }
}
