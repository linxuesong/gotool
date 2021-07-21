package linkedlist

import (
	"fmt"
	"testing"
)

func TestSingleLinkedList_Reverse(t *testing.T) {
	lst := NewSingleLinkedList()
	lst.AddintoHead(4)
	lst.AddintoHead(3)
	lst.AddintoHead(2)
	fmt.Printf("%#v \n", lst.ToString())
	lst.Reverse()
	fmt.Printf("%#v \n", lst.ToString())
}