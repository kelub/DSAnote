package BackTracking

import (
	"fmt"
)

type Queens struct {
	result  []int   // 存放每次递归结果，下标代表行值代表当前行的列
	num     int     // 棋盘大小
	results [][]int // 存放所有解
}

func (q *Queens) queens(row int) {
	if row == q.num {
		fmt.Println("xxx", q.result)
		r := make([]int, q.num, q.num)
		copy(r, q.result)
		q.results = append(q.results, r)
		return
	}
	for column := 0; column < q.num; column++ {
		if q.check(row, column) {
			fmt.Println("row, column", row, column)
			q.result[row] = column
			q.queens(row + 1)
		}
	}
}

func (q *Queens) check(row, column int) bool {
	//左对角线
	leftup := column - 1
	//右对角线
	rightup := column + 1
	for i := row - 1; i >= 0; i-- {
		if q.result[i] == column {
			return false
		}
		if leftup >= 0 {
			if q.result[i] == leftup {
				return false
			}
		}
		if rightup <= q.num-1 {
			if q.result[i] == rightup {
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
	items  []int // 每个物品重量数组
	weight int   // 物品总总量

	maxWeight int // 最大容量
}

// do
// curw 当前总总量
// weight 物品总总量
// n 物品个数
// items 当前物品重量
func (p *BackPack) do(i, curw, weight, n int, items []int) {
	if curw >= weight || i == n { // 背包满或者物品装完 结束
		if p.maxWeight < curw {
			p.maxWeight = curw
		}
		return
	}
	// 不装
	p.do(i+1, curw, weight, n, items)
	// 装
	if weight >= curw+items[i] {
		p.do(i+1, curw+items[i], weight, n, items)
	}
}

type Pattern struct {
	matched bool
}

func (p *Pattern) match(text, pattern []rune) bool {
	p.matched = false
	p.do(0, 0, text, pattern, len(pattern), len(text))
	return p.matched
}

// jsa19ks  j*k?s
// do
// text ,pattern 匹配字符串，正则表达式
// plen 正则长度  pindex 当前正则表达式位置 textindex 匹配字符串当前位置
func (p *Pattern) do(pindex, textindex int, text, pattern []rune, plen, tlen int) {
	if pindex == plen { // 正则表达式结束退出
		if textindex == tlen { // 匹配字符串到头代表匹配成功
			p.matched = true
		}
		return
	}
	// 匹配任意个字符
	if string(pattern[pindex]) == "*" {
		for i := textindex; i < tlen; i++ {
			p.do(pindex+1, i, text, pattern, plen, tlen)
		}
	} else if string(pattern[pindex]) == "?" { // 匹配一个字符或者空
		// 匹配0个字符
		p.do(pindex+1, textindex, text, pattern, plen, tlen)
		// 匹配一个字符
		p.do(pindex+1, textindex+1, text, pattern, plen, tlen)
	} else if textindex < tlen && pattern[pindex] == text[textindex] { // 匹配纯字符串
		p.do(pindex+1, textindex+1, text, pattern, plen, tlen)
	}
}

//分割回文串
func Partition(s string) [][]string {
	p := &P{
		slen:    len(s),
		results: make([][]string, 0),
		mem:     make([][]bool, len(s)+1),
	}
	result := make([]string, 0)
	for i, _ := range p.mem {
		k := make([]bool, len(s)+1)
		p.mem[i] = k
	}
	p.spilt(s, 0, result)
	return p.results
}

type P struct {
	slen    int
	results [][]string

	// mem[i][j] 值代表 从i-j 是否是回文
	mem [][]bool

	test int
}

func (p *P) spilt(s string, index int, res []string) {
	result := make([]string, len(res))
	copy(result, res)
	if index >= p.slen {
		p.results = append(p.results, result)
		return
	}

	for i := index + 1; i <= p.slen; i++ {
		if !p.mem[index][i] {
			if p.isPalindrome(s[index:i]) {
				p.mem[index][i] = true
				r := append(result, s[index:i])
				p.spilt(s, i, r)
				if i > 10 {
					return
				}
			}
		} else {
			r := append(result, s[index:i])
			p.spilt(s, i, r)
			if i > 10 {
				return
			}
		}
	}
	return
}

func (p *P) isPalindrome(s string) bool {
	b := []byte(s)
	slen := len(b)
	head, end := 0, slen-1
	for head < end {
		if b[head] == b[end] {
			head++
			end--
			continue
		} else {
			return false
		}
	}
	return true
}
