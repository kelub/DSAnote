package BackTracking

import (
	"fmt"
)

type Queens struct {
	result []int    // 存放每次递归结果，下标代表行值代表当前行的列
	num 	int 	// 棋盘大小
	results [][]int // 存放所有解
}

func (q *Queens) queens(row int){
	if row == q.num{
		fmt.Println("xxx",q.result)
		r := make([]int,q.num,q.num)
		copy(r,q.result)
		q.results = append(q.results,r)
		return
	}
	for column := 0;column<q.num;column++{
		if q.check(row,column){
			fmt.Println("row, column",row,column)
			q.result[row] = column
			q.queens(row+1)
		}
	}
}

func (q *Queens)check(row ,column int)bool{
	//左对角线
	leftup := column - 1
	//右对角线
	rightup := column + 1
	for i := row-1;i>=0;i--{
		if q.result[i] == column{
			return false
		}
		if leftup >= 0{
			if q.result[i] == leftup{
				return false
			}
		}
		if rightup <= q.num-1{
			if q.result[i] == rightup{
				return false
			}
		}
		// 上移行 更新对角线
		leftup--
		rightup++
	}
	return true
}


// 1,2,5,10,15,20,25,50
type BackPack struct {
	items []int 		// 每个物品重量数组
	weight	int			// 物品总总量

	maxWeight 	int  	// 最大容量
}

// do
// curw 当前总总量
// weight 物品总总量
// n 物品个数
// items 当前物品重量
func (p *BackPack)do(i,curw ,weight,n int, items []int){
	if curw >= weight || i == n{ // 背包满或者物品装完 结束
		if p.maxWeight < curw{
			p.maxWeight = curw
		}
		return
	}
	// 不装
	p.do(i+1,curw,weight,n,items)
	// 装
	if weight >= curw+items[i]{
		p.do(i+1,curw+items[i],weight,n,items)
	}
}