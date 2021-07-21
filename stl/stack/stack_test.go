package stack

import (
	"fmt"
	"testing"
)

func TestStack_Push(t *testing.T) {
	sta := NewStack()
	sta.Push(1)
	sta.PrintStack()
	sta.Push(2)
	sta.PrintStack()
	sta.Push(3)
	sta.PrintStack()
	sta.Push(4)
	sta.PrintStack()
	sta.Pop()
	sta.PrintStack()
	sta.Pop()
	sta.PrintStack()
	sta.Pop()
	sta.PrintStack()
	fmt.Printf("%v",sta.Pop())
	sta.PrintStack()
	fmt.Printf("%v",sta.Pop())
	sta.PrintStack()
	fmt.Printf("%v",sta.Pop())
	sta.PrintStack()

}
