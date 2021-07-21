package math

import (
	"testing"
)

func TestMedian(t *testing.T) {
	nums := []int{2,1,4,3,2,5}
	t.Logf("before sort nums:%v", nums)
	QuickSort(nums, 0, len(nums)-1)
	t.Logf("after sort nums:%v", nums)
	t.Logf("中位数:%v", Median(nums))
}

func TestSwap(t *testing.T) {
	nums := make([]int, 2)
	nums[0],nums[1] = 1,2

	Swap(&nums[0], &nums[1])
	t.Logf("a:%v, b:%v", nums[0], nums[1])
}

func TestQuickSort(t *testing.T) {
	nums := []int{3,1,5,4,3,2}
	t.Logf("before sort nums:%v", nums)
	QuickSort(nums, 0, len(nums)-1)
	t.Logf("after sort nums:%v", nums)
}

func TestOneQuickSort(t *testing.T) {
	nums := []int{3,1,5,4,3,2}
	t.Logf("before sort nums:%v", nums)
	t.Logf("OneQuickSort return %v, after one sort nums:%#v", OneQuickSort(nums, 0, len(nums)-1), nums)
}

func TestOneQuickSort2(t *testing.T) {
	nums := []int{3,1,5,4,3,2}
	t.Logf("before sort nums:%v", nums)
	t.Logf("OneQuickSort return %v, after one sort nums:%#v", OneQuickSort2(nums, 0, len(nums)-1), nums)
}

func TestTopKNum(t *testing.T) {
	nums := []int{3,1,5,4,2}
	t.Logf("top k max num:%v", TopKMinNum(nums,0,len(nums)-1, 3))
}