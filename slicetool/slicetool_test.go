package slicetool

import (
	"fmt"
	"testing"
)

func TestUpSetSlice(t *testing.T) {
	arr := []int{1, 2, 3, 4}
	ShuffleSlice(arr)
	fmt.Printf("%#v", arr)
}
