package queue

import (
	"errors"
	"fmt"
	"github.com/linxuesong/gotool/stl/linkedlist"
)

type Queue struct {
	 Head *linkedlist.SingleLinkedListNode
	 Tail *linkedlist.SingleLinkedListNode
	 Size int
}

func NewQueue() Queue{
	return Queue{
		Head: nil,
		Tail: nil,
		Size: 0,
	}
}

// 尾插, 为什么不头插入尾出，尾出时间复杂度为O(N)
func(this *Queue) Enqueue(val interface{}) {
	inode := linkedlist.SingleLinkedListNode{
		Item: val,
		Next: nil,
	}
	if this.Tail == nil {
		this.Head = &inode
		this.Tail = &inode
	}else {
		this.Tail.Next = &inode
		this.Tail = &inode
	}

	this.Size++
}
// 头出
func(this *Queue) Dequeue() error{
	if this.Size == 0 {
		return errors.New("error, try dequeue an empty Queue")
	}
	this.Head = this.Head.Next
	this.Size--
	return nil
}

func(this Queue) PrintQueue() {
	fmt.Printf("Queue content font<-tail: ")
	for cur := this.Head; cur != nil; cur = cur.Next {
		fmt.Printf("<- %v ", cur)
	}
	fmt.Println()
}