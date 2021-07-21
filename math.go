package gotool

import "github.com/linxuesong/gotool/math"

func Medium(arr []int) float64 {
	return math.Median(arr)
}

func Sort(arr []int) {
	arr_len := len(arr)
	if arr_len<=1 {
		return
	}
	math.QuickSort(arr, 0, arr_len-1)
}