package BinarySearch

import (
	"errors"
	"fmt"
	"math"
)

type BSSlice struct {
	a     []int
	value int
}

func (bs *BSSlice) Do() {
	bs.Iterator(bs.a, bs.value)
}

func (bs *BSSlice) Iterator(a []int, value int) (int, error) {
	if len(a) == 0 {
		return -1, errors.New("nil")
	}
	alen := len(a)
	var (
		low  = 0
		high = alen
	)
	// low >= high
	for low <= high {
		//
		mid := low + ((high - low) >> 1)
		if value == a[mid] {
			return mid, nil
		}
		if a[mid] > value {
			//
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1, errors.New("not exist")
}

func (bs *BSSlice) Sqrt(value float64) (float64, error) {
	if value == 0 {
		return -1, errors.New("value is 0")
	}
	var (
		low  = float64(0)
		high = value
		mid  = value / 2
	)
	//similar := mid * mid
	for similar := mid * mid; math.Abs(similar-value) > 0.000001; similar = mid * mid {
		//for math.Abs(similar-value) > 0.000001 {
		fmt.Println(similar)
		if similar == value {
			return mid, nil
		}
		if similar > value {
			high = mid
		} else {
			low = mid
		}
		mid = low + ((high - low) / 2)
		//similar = mid * mid
	}
	return mid, nil
}

// A 查找第一个值等于给定值的元素
// B 查找最后⼀个值等于给定值的元素
// C 查找第⼀个⼤于等于给定值的元素
// D 查找最后⼀个⼩于等于给定值的元素

// A
func (bs *BSSlice) FirstEqualValue(a []int, value int) int {
	if len(a) == 0 {
		fmt.Println("nil")
		return -1
	}
	alen := len(a)
	var (
		low  = 0
		high = alen
	)
	for low <= high {
		mid := low + ((high - low) >> 1)

		if value == a[mid] {
			//return bs.first(a, value, mid) // 当重复数过多时候复杂度会上升,最坏变成 O(n) 不符合 二分思想
			// 应当继续二分，如果满足边界条件，直接返回 index
			if mid == 0 || a[mid-1] < value {
				return mid
			} else {
				high = mid - 1
			}
		}

		if a[mid] > value {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	fmt.Println("not exist")
	return -1
}

// B
func (bs *BSSlice) LastEqualValue(a []int, value int) int {
	if len(a) == 0 {
		fmt.Println("nil")
		return -1
	}
	alen := len(a)
	var (
		low  = 0
		high = alen - 1
	)
	for low <= high {
		mid := low + ((high - low) >> 1)

		if value == a[mid] {
			//return bs.last(a, value, mid) // 当重复数过多时候复杂度会上升,最坏变成 O(n) 不符合 二分思想
			// 应当继续二分，如果满足边界条件，直接返回 index
			if mid == alen-1 || a[mid+1] > value {
				return mid
			} else {
				low = mid + 1
			}
		}

		if a[mid] > value {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	fmt.Println("not exist")
	return -1
}

// A
func (bs *BSSlice) first(a []int, value int, mid int) int {
	r := 0
	for i := mid - 1; i >= 0; i-- {
		if a[i] < value {
			r = i + 1
			return r
		}
	}
	return r
}

// B
func (bs *BSSlice) last(a []int, value int, mid int) int {
	r := len(a) - 1
	for i := mid + 1; i < len(a); i++ {
		if a[i] > value {
			r = i - 1
			return r
		}
	}
	return r
}

// C
func (bs *BSSlice) FirstGreaterEqualValue(a []int, value int) int {
	if len(a) == 0 {
		fmt.Println("nil")
		return -1
	}
	alen := len(a)
	var (
		low  = 0
		high = alen
	)
	for low <= high {
		mid := low + ((high - low) >> 1)
		if value <= a[mid] {
			if mid == 0 || a[mid-1] < value {
				return mid
			} else {
				high = mid - 1
			}
		} else {
			low = mid + 1
		}
	}
	fmt.Println("not exist")
	return -1
}

// D
func (bs *BSSlice) LastLessEqualValue(a []int, value int) int {
	if len(a) == 0 {
		fmt.Println("nil")
		return -1
	}
	alen := len(a)
	var (
		low  = 0
		high = alen - 1
	)
	for low <= high {
		mid := low + ((high - low) >> 1)
		if value >= a[mid] {
			if mid == alen-1 || a[mid+1] > value {
				return mid
			} else {
				low = mid + 1
			}
		} else {
			high = mid - 1
		}
	}
	fmt.Println("not exist")
	return -1
}

// 循环有序数 4，5，6，1，2，3
// 4 4 1 2 3    1 2 3 4 5 6
// A

func (bs *BSSlice) FirstEqualValueRing(a []int, value int) int {
	alen := len(a)
	//确定分界
	var inter int = alen - 1
	for i := 0; i < alen-1; i++ {
		if a[i] > a[i+1] {
			inter = i
		}
	}
	if inter == alen-1 {
		return bs.FirstEqualValue(a, value)
	}

	if value < a[alen-1] {
		return bs.FirstEqualValue(a[inter+1:], value) + inter
	} else if value > a[0] {
		return bs.FirstEqualValue(a[:inter], value)
	} else if value == a[alen-1] {
		return alen - 1
	} else if value == a[0] {
		return 0
	} else {
		return -1
	}
	return -1
}
