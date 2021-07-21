package stack

import (
	"fmt"
	"github.com/linxuesong/gotool/stl/linkedlist"
)

type Stack struct{
	lst *linkedlist.SingleLinkedList
	size int
}

func NewStack() Stack{
	return Stack{
		lst: linkedlist.NewSingleLinkedList(),
        size: 0,
	}
}

func(this *Stack) Pop() error {
	err := this.lst.DeletefromHead()
	if err != nil {
		return err
	}
	this.size--
	return nil
}

func(this *Stack) Push(val interface{}) {
	this.lst.AddintoHead(val)
	this.size++
}

func(this Stack) IsEmpty() bool{
	return this.size==0
}

func(this Stack) PrintStack() {
	fmt.Printf("Stack content top->bottom: ")
	for cur := this.lst.Head; cur != nil; cur = cur.Next {
		fmt.Printf("-> %v ", cur)
	}
	fmt.Println()
}
