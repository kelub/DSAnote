package Sort

//递推公式
/*
p-->r  q 为分割点
quick_sort(p,r) = quick_sort(p,q-1) + quick_sort(q+1,r)
终止条件
p>=r
关键点
获取分割点q并使得 p,q-1 < q < q+1,r
*/

func QuickSort(a []int) {
	quicksort(a, 0, len(a)-1)
}

func quicksort(a []int, p int, r int) {
	if p >= r {
		return
	}
	q := pivot(a, p, r)
	return
	//left
	quicksort(a, p, q-1)
	//right
	quicksort(a, q+1, r)
}

func pivot(a []int, q int, r int) (p int) {
	//取点
	pivot := a[r]
	p = q
	//q,p 已处理区间 全部小于 pivot
	//p+1,r 未处理区间
	for j := q; j < r; j++ {
		if a[j] < pivot {
			a[p], a[j] = a[j], a[p]
			p = p + 1
		}
	}
	a[p], a[r] = a[r], a[p]
	return p
}
