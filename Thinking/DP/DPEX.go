package DP

import (
	"fmt"
	//"fmt"
	"math"
	"math/rand"
)

/*
n乘以n的矩阵w[n][n]。
将棋⼦从左上⻆移动到右下⻆。每次只能向右或者向下移动⼀位。从左上⻆到右下⻆.
求最短路径

*/
type DistBT struct {
	n       int
	minDist int
	base    [][]int

	layer int

	// mem 备忘录 i,j 值存放当前最小路径值
	mem [][]int

	// dp
	states [][]int
}

func NewDistBT(n int) *DistBT {
	base := make([][]int, n, n)
	d := &DistBT{
		n:       n,
		minDist: math.MaxInt32,
		layer:   (2 * n) - 1,
	}

	for i := 0; i < n; i++ {
		base[i] = make([]int, n, n)
		for j := 0; j < n; j++ {
			base[i][j] = rand.Intn(10)
		}
	}
	d.base = base
	return d
}

// 暴力1  存在重复子问题
// (0,0,0,d.base[0][0])
func (d *DistBT) do1(layer int, i, j int, curDist int) {
	if layer == d.layer-1 {
		if d.minDist > curDist {
			d.minDist = curDist
		}
		return
	}
	// up
	if i < d.n-1 {
		d.do1(layer+1, i+1, j, curDist+d.base[i+1][j])
	}
	// right
	if j < d.n-1 {
		d.do1(layer+1, i, j+1, curDist+d.base[i][j+1])
	}
}

// mem 备忘录版本  mem二维数组存放路径，并更新最小路径
// 重复点在于 i,j 重复。当计算到当前点有最小值时，
// 递归回溯的其它路径重复,或者无效路径
func (d *DistBT) domem(layer int, i, j int, curDist int) {
	if layer == d.layer-1 {
		if d.minDist > curDist {
			d.minDist = curDist
		}
		return
	}
	// 去除重复无效递归路径
	if d.mem[i][j] != -1 && d.mem[i][j] < curDist {
		return
	}

	d.mem[i][j] = curDist
	// up
	if i < d.n-1 {
		d.do1(layer+1, i+1, j, curDist+d.base[i+1][j])
	}
	// right
	if j < d.n-1 {
		d.do1(layer+1, i, j+1, curDist+d.base[i][j+1])
	}
}

// 动态规划
// 利用数组记录，当前状态由上一个状态通过数组而来。
// curMinDist = curValue + min(pre(up),pre(right))
// minDist(i,j) = d.base[i][j] + min(minDist(i-1,j),minDist(i,j-1))
// 这个状态表示恰好和暴力回溯(递归)相反，也和动态规划递归版相反
// 状态转移表
func (d *DistBT) dp() int {
	// init
	d.states[0][0] = d.base[0][0]
	for i := 0; i < d.n; i++ {
		for j := 0; j < d.n; j++ {
			var preMinDist int
			if i == 0 && j == 0 {
				preMinDist = 0
			} else if i == 0 && j > 0 {
				preMinDist = d.states[0][j-1]
			} else if i > 0 && j == 0 {
				preMinDist = d.states[i-1][0]
			} else {
				preMinDist = int(math.Min(float64(d.states[i-1][j]), float64(d.states[i][j-1])))
			}
			d.states[i][j] = d.base[i][j] + preMinDist
		}
	}
	d.minDist = d.states[d.n-1][d.n-1]
	return d.states[d.n-1][d.n-1]
}

// 状态转移方程  递归解法 和迭代解法相反
// 迭代解法用状态数组这个视角配合状态方程，更容易理解。
// 递归解法用状态方程这个视角配合迭代加上备忘录，相对不太好理解，但是实现起来更简洁。
// curMinDist 为当前 i,j状态的最小路径值，
// dpdo(d.n-1,d.n-1) curMinDist
func (d *DistBT) dpdo(i, j int) (curMinDist int) {
	// 去除重复分支--剪枝
	if d.mem[i][j] > 0 {
		return d.mem[i][j]
	}
	// 终止条件
	if i == 0 && j == 0 {
		curMinDist = d.base[0][0]
		return
	} else if i > 0 && j == 0 {
		curMinDist = d.base[i][j] + d.dpdo(i-1, j)
	} else if i == 0 && j > 0 {
		curMinDist = d.base[i][j] + d.dpdo(i, j-1)
	} else {
		curMinDist = d.base[i][j] + int(math.Min(float64(d.dpdo(i-1, j)), float64(d.dpdo(i, j-1))))
	}
	// 记录i,j 这个状态的最小路径
	d.mem[i][j] = curMinDist
	return curMinDist
}

//// 最长子串
//
///*
//状态
//	i,j a,b 指针
//	max_lcs 最长公共子串数
//回溯算法(直观从左到右迭代)
//if a[i] == b[j]
//	max_lcs = max_lcs +1 迭代 i+1,j+1
//if a[i] != b[j]
//	max_lcs = max_lcs 迭代 i,j+1 或者 迭代 i+1,j
//
//转移方程(迭代方式，从最后结果反推)
//if a[i] == b[j]
//	max_lcs(i,j) = max(max_lcs(i-1,j-1)+1, max_lcs(i-1,j), max_lcs(i,j-1))
//if a[i] != b[j]
//	max_lcs(i,j) = max(max_lcs(i-1,j-1), max_lcs(i-1,j), max_lcs(i,j-1))
//*/
//type MaxSub struct {
//	states [][]int
//}
//
//func (m *MaxSub) maxSub(a, b string, max int) {
//
//}

// 最⻓递增⼦序列⻓度
/*
2, 9, 3, 6, 5, 1, 7
2, 3, 5, 7
4

a[0...i] 的最大子序列长度
maxLen(i) = 比a[i]小的a[j] maxLen(j) + 1
画图增加理解 继续具象化 状态转移方程

maxLen(i) = max(for j<i (maxLen(j)+(if a[i]>a[j] 1  else 0 )))
*/

func MaxIncSubf(a []byte) {
	m := &MaxIncSub{
		state: make([][]int, len(a)),
	}
	// init
	for i := 0; i < len(a); i++ {
		m.state[i] = make([]int, len(a))
	}
	m.state[0][0] = 0
	m.state[0][1] = 2
	m.maxIncSub(a)
	fmt.Println("state", m.state[len(a)-2][len(a)-1])
}

type MaxIncSub struct {
	state [][]int
}

// 迭代方式(根据状态转移方程填充DB table)
func (m *MaxIncSub) maxIncSub(a []byte) {
	for j := 2; j < len(a); j++ {
		var max = 0
		for i := j - 1; i >= 0; i-- {
			if a[j] > a[i] {
				if i == 0 {
					m.state[i][j] = 1 + 1
				} else {
					m.state[i][j] = m.state[i-1][j-1] + 1
				}
			} else {
				if i == 0 {
					m.state[i][j] = 1
				} else {
					m.state[i][j] = m.state[i-1][j-1]
				}
			}
			if max < m.state[i][j] {
				max = m.state[i][j]
			}
		}
		m.state[j-1][j] = max
	}
}
