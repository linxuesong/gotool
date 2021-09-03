package math

/**
 * @Author linxuesong
 * @Description 中位数,三种实现方法  1.排序 2.大顶堆 3.快排找第n/2大的数   这里使用方式3//TODO
 * @Date 8:16 下午 2021/7/12
 * @Param  * @Param: null
 * @return
 **/
func Median(nums []int) float64 {
	length := len(nums)
	pre := TopKMinNum(nums, 0, length-1, (length)/2) // 第n/2小的数

	if length%2 == 1 {
		return float64(pre)
	}
	cur := TopKMinNum(nums, 0, length-1, (length)/2-1) // 第n/2-1小的数
	return float64(pre+cur) / 2
}

// 第k小的数, k为0标识最小的数
func TopKMinNum(nums []int, sta, end, k int) int {
	site := OneQuickSort(nums, sta, end)
	if site == k {
		return nums[k]
	} else if site > k {
		return TopKMinNum(nums, sta, site-1, k)
	} else {
		return TopKMinNum(nums, site+1, end, k)
	}
}

func QuickSort(nums []int, sta, end int) {
	if sta >= end {
		return
	}
	site := OneQuickSort(nums, sta, end)
	QuickSort(nums, sta, site-1)
	QuickSort(nums, site+1, end)
}

func OneQuickSort(nums []int, sta, end int) int {
	mark, first := sta+1, nums[sta] //  mark左开区间均小于等于基准     [s,mark)小于基准    [mark,+)大于基准

	for i := sta + 1; i <= end; i++ {
		if first >= nums[i] { // 符合mark标记元素,加入到mark中
			Swap(&nums[mark], &nums[i])
			mark++
		}
	}
	// mark左侧都是小于等于基准的元素,但是基准在第一个位置           num[mark]大于基准
	Swap(&nums[sta], &nums[mark-1]) // 把基准元素放到恰当位置上
	return mark - 1
}

func OneQuickSort2(nums []int, sta, end int) int {
	lo, hi := sta+1, end
	for lo <= hi {
		for lo <= hi && nums[lo] <= nums[sta] {
			lo++
		}
		for lo <= hi && nums[hi] >= nums[sta] {
			hi--
		}
		Swap(&nums[lo], &nums[hi])
	}
	Swap(&nums[sta], &nums[lo-1])
	return lo - 1
}

func Swap(a, b *int) {
	c := *b
	*b = *a
	*a = c
}

/**
 * @Author linxuesong
 * @Description  search and return the index of key in nums or -1 when key is not existed in nums
 *
 * @Data files:   https://algs4.cs.princeton.edu/11model/tinyAllowlist.txt
 *                https://algs4.cs.princeton.edu/11model/tinyText.txt
 *                https://algs4.cs.princeton.edu/11model/largeAllowlist.txt
 *                https://algs4.cs.princeton.edu/11model/largeText.txt
 * @Date 8:08 下午 2021/7/14
 * @Param nums the array of int, must be sorted in ascending ordier
 * @Param key is the taget that search
 * @return index of key in arr if persent;otherwise -1 if absent
 **/
func BinarySearch(nums []int, key int) int {
	lo, hi := 0, len(nums)-1
	for lo <= hi {
		mid := (lo + hi) / 2
		if key == nums[mid] {
			return mid
		} else if key > nums[mid] {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return -1
}
