package Sort

/*
递推公式

merge_sort(a,q,r) = merge(merge_sort(a,q,(q+r)/2) + merge_sort(a,(q+r)/2-1 ,r))

另一种表示

merge_sort(a,q,r) = merge_sort(a,q,(q+r)/2) + merge_sort(a,((q+r)/2) - 1,r) + merge(a,q,r,(q+r)/2)

终止条件
q >= r  无法再分

关键点
合并两个已经排好序数组，形成新的已排好序数组。
*/

/*
时间复杂度(递归) O(nlogn)
空间复杂度 n
由于 merge 都是递归后再 for 对比，平均，最好，最坏都一样。
稳定排序
空间复杂度 merge 需要一个 n/k 的空间存储临时排序好的数据。相同时刻只有一个函数运行，
所以不会累加。
递归树分析法(直观)：
		n
		n/2 + n/2
		n/k ...+...n/k
单个 merge 时间复杂度为 O(n/k) k 为递归二分，也就是当前 merge 递归树深度。
整个复杂度为全部 merge 相加。
复杂度等于每个节点相加。恰好每层节点相加 =n 树高度 =log(n)

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
