package Greedy

import(
	"testing"
	"fmt"
)

/*
[6,8] [2,4] [3,5] [1,5] [5,9] [8,10]
不相交区间
[2,4] [6,8] [8,10]
*/
func TestInterval(t *testing.T){
	a := [][]int{[]int{6,8}, []int{2,4}, []int{3,5}, []int{1,5}, []int{5,9}, []int{8,10}}
	result := make([][]int,0)
	Interval(a,&result)
	fmt.Println("r",result)
}