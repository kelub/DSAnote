package Greedy

import "fmt"


func Interval(a [][]int,result *[][]int){
	lmin ,lmax := 0,0
	for i := range a{
		if lmin > a[i][0]{
			lmin = a[i][0]
		}
		if lmax < a[i][1]{
			lmax = a[i][1]
		}
	}
	// 排序a
	QuickSort(a)
	fmt.Println("a",a)
	do(a,0,len(a),lmin ,lmax,result)
}

func do(a [][]int,q,r int ,lmin ,lmax int,result *[][]int) {
	if q > r -1{
		//fmt.Println(result)
		return
	}
	k := false
	min := a[q][1]
	k = true
	item := q
	for i := q;i<r;i++{
		if a[i][0] < lmin{
			min = a[i+1][1]
			k = false
			continue
		}
		if a[i][1] < min{
			min = a[i][1]
			item = i
			k = true
		}
	}
	if k{
		*result = append(*result,a[item])
		do(a,item +1,r,min ,lmax,result)
	}
}

func QuickSort(a [][]int){
	quicksort(a,0,len(a)-1)
}

func quicksort(a [][]int,q ,r int){
	if q >= r{
		return
	}
	p := pivot(a,q,r)
	quicksort(a,q,p-1)
	quicksort(a,p+1,r)

}

func pivot(a [][]int,q,r int)int{
	pivot := a[r][0]
	p := q
	// q.p 已处理区间，即全部小于 pivot 的值
	// p+1,r 未处理区间
	for i := q;i<r;i++{
		if a[i][0] < pivot{
			a[p],a[i] = a[i],a[p]
			p++
		}
	}
	a[p],a[r] = a[r],a[p]
	return p
}