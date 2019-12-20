package Recursive

import (
	"fmt"
)

//全排列

/*
1 2 3 4 5
递推公式
f(n) = (最后一位是 1 + f(n-1)) + (最后一位是 2 + f(n-1)) + ... + (最后一位是 n + f(n-1))
*/

/*
时间复杂度(递归树)
每层	总数
n		n
n-1		n(n-1)
n-2		n(n-1)(n-2)
...		...
1		n!
*/

func Fully(a []int, n int, k int) {
	// fmt.Printf("%d ", k)
	if k == 1 {
		for i := 0; i < n; i++ {
			fmt.Printf("%d ", a[i])
		}
		fmt.Println("")
	}
	for i := 0; i < k; i++ {
		// 交换
		temp := a[i]
		a[i] = a[k-1]
		a[k-1] = temp
		Fully(a, n, k-1)
		temp = a[i]
		a[i] = a[k-1]
		a[k-1] = temp
	}
}
