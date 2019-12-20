package Divide

/*
2 4 3 1 5 6
*/

var result = 0
func ReverseOrder(a []int){
	merge_sort(a,0,len(a)-1)
}

func merge_sort(a []int,q,r int){
	if q >= r{
		return
	}
	mid := (q + r)/2
	merge_sort(a,q,mid)
	merge_sort(a,mid +1,r)
	merge(a,q,mid,r)
}

// 2 3 4   1 5 6
func merge(a []int,q,mid,r int){
	i,j,k := q,mid + 1,0
	temp := make([]int,r-q+1,r-q+1)
	for i<=mid && j<=r{
		if a[i]<a[j]{
			temp[k] = a[i]
			i++
		}else{
			temp[k] = a[j]
			j++
			result += mid-i+1
		}
		k++
	}
	for i<=mid{
		temp[k] = a[i]
		k++
		i++
	}
	for j<=r{
		temp[k] = a[j]
		k++
		j++
	}

	for i := 0;i<r-q+1;i++{
		a[q+i] = temp[i]
	}
}
