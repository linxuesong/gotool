package slicetool

import "math/rand"

// 注意:rand库已经提供了随机打乱数组的方法 rand.Shuffle
// 打乱数组, 使用洗牌算法 将第i张排和后面i-1章牌随机交换， i为下标从n-1利到1
func ShuffleSlice(arr []int) {
	leng := len(arr)
	for i := leng - 1; i >= 1; i-- {
		ri := rand.Intn(i)
		c := arr[ri]
		arr[ri] = arr[i]
		arr[i] = c
	}
}
