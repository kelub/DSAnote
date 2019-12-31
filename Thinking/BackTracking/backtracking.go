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
func Partition(s string) *P {
	p := &P{
		slen:    len(s),
		result:  make([][]byte, 0),
		results: make([][][]byte, 0),
	}
	for i := 1; i < p.slen; i++ {
		p.spilt([]byte(s), i)
	}
	return p
}

type P struct {
	slen    int
	result  [][]byte
	results [][][]byte
}

func (p *P) spilt(s []byte, indes int) {
	if indes >= p.slen-1 {
		temp := make([][]byte, len(p.result))
		copy(temp, p.result)
		p.results = append(p.results, temp)
		p.result = make([][]byte, 0)
		return
	}
	fmt.Println("s", s[0:indes], !p.isPalindrome(s[0:indes]))
	if !p.isPalindrome(s[0:indes]) {
		return
	}
	p.result = append(p.result, s[0:indes])

	for i := 1; i < len(s[:indes]); i++ {
		p.spilt(s[indes:], i)
	}
}

func (p *P) isPalindrome(s []byte) bool {
	slen := len(s)
	head, end := 0, slen-1
	for head < end {
		if s[head] == s[end] {
			head++
			end--
			continue
		} else {
			return false
		}
	}
	return true
}
