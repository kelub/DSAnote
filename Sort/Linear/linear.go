package Linear

type LinearSort struct {
}

func New() *LinearSort {
	return &LinearSort{}
}

// 3 6 2 2 0 4 0 5
func (ls *LinearSort) CountingSort(a []int) bool {
	if a == nil || len(a) == 0 {
		return false
	}

	// 得到分桶数
	max := a[0]
	for i := 1; i < len(a); i++ {
		if max < a[i] {
			max = a[i]
		}
	}

	// 创建记数数组
	count := make([]int, max+1, max+1)

	// 得到计数值
	for i := 0; i < len(a); i++ {
		count[a[i]]++
	}

	// 累加计数值
	for i := 1; i < len(count); i++ {
		count[i] = count[i-1] + count[i]
	}

	// 计数排序关键
	// 计数数组通过累加方式，记录数据的位置
	r := make([]int, len(a), len(a))

	for i := 0; i < len(a); i++ {
		index := count[a[i]] - 1
		r[index] = a[i]
		count[a[i]]--
	}

	// copy
	for i := 0; i < len(a); i++ {
		a[i] = r[i]
	}
	return true
}

/*
a  	3 6 2 2 0 4 0 5
c	2 0 2 1 1 1 1
c	2 2 4 5 6 7 8
*/
