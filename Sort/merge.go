package Sort

/*
递推公式

merge_sort(a,q,r) = merge(merge_sort(a,q,(q+r)/2) + merge_sort(a,(q+r)/2-1 ,r))

终止条件
q >= r  无法再分

关键点
合并两个已经排好序数组，形成新的已排好序数组。
*/

func MergeSort(a []int) {
	merge_sort(a, 0, len(a)-1)
}

func merge_sort(a []int, q int, r int) {
	if q >= r {
		return
	}
	min := (q + r) / 2
	// left
	merge_sort(a, q, min)
	// right
	merge_sort(a, min+1, r)
	//left := a[q:min]
	//right := a[min+1:r]
	merge(a, q, r, min)
}

func merge(a []int, q, r, min int) {
	temp := make([]int, r-q+1)
	i, j, k := q, min+1, 0
	for i <= min && j <= r {
		if a[i] < a[j] {
			temp[k] = a[i]
			i++
		} else {
			temp[k] = a[j]
			j++
		}
		k++
	}
	//left剩余放入temp
	for i <= min {
		temp[k] = a[i]
		i++
		k++
	}
	// right剩余放入temp
	for j <= r {
		temp[k] = a[j]
		j++
		k++
	}

	for i := 0; i < r-q+1; i++ {
		a[q+i] = temp[i]
	}
}
