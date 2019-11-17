package BackTracking

import(
	"testing"
	"fmt"
)

func TestQueens(t *testing.T){
	num := 4
	q := Queens{
		result:make([]int,num,num),
		num:num,
		results:make([][]int,0),
	}
	q.queens(0)
	fmt.Println(q.results)
}

func TestBackPack(t *testing.T){
	p := BackPack{
		items: []int{1,2,4,8,16,32},
		weight:100,
	}
	p.do(0,0,p.weight,len(p.items),p.items)
	fmt.Println(p.maxWeight)
}
