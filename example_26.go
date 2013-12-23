/*
 * linked list implementation
 */

 package main

 import (
    "fmt"
 )

 type Node struct {
    next *Node
    data int
}

var head *Node=nil
func addNode(x int) {

        node:=&Node{} //when assigning the structure use '&'
        node.next=head
        node.data=x
        head=node
        fmt.Println(head)
        fmt.Println(head.next)


}

func delNode(x1 int) {

    y:=&Node{}
    y.next=head
    for x:=head;x.next!=nil;x=x.next {
        if x1==x.data {
            y.next=head.next
            head.next=nil
            y.next=head
        }
    }
}
func display() {
    fmt.Println(head)
    fmt.Println(head.next)
    for x:=head;x.next!=nil;x=x.next {
        fmt.Println(x.next)
        fmt.Println(x.data)
    }
}

func main() {
    addNode(10)
    addNode(2)
    addNode(6)
    display()
    delNode(2)

    display()
}
