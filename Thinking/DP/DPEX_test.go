package DP

import (
	"fmt"
	"testing"
)

func Test_DistBT(t *testing.T) {
	d := NewDistBT(4)

	fmt.Println("base", d.base)
	d.do1(0, 0, 0, d.base[0][0])
	fmt.Println(d.minDist)
}

func Test_DistBT_DOMeM(t *testing.T) {
	d := NewDistBT(4)

	fmt.Println("base", d.base)
	mem := make([][]int, d.n, d.n)
	for i := 0; i < d.n; i++ {
		mem[i] = make([]int, d.n, d.n)
		for j := 0; j < d.n; j++ {
			mem[i][j] = -1
		}
	}
	d.mem = mem
	d.domem(0, 0, 0, d.base[0][0])
	fmt.Println(d.minDist)
	fmt.Println(d.mem)
}

func Test_DistBT_DP(t *testing.T) {
	d := NewDistBT(4)

	fmt.Println("base", d.base)
	mem := make([][]int, d.n, d.n)
	for i := 0; i < d.n; i++ {
		mem[i] = make([]int, d.n, d.n)
		for j := 0; j < d.n; j++ {
			mem[i][j] = -1
		}
	}
	d.states = mem
	d.dp()
	fmt.Println(d.minDist)
	fmt.Println(d.states)
}

func Test_DistBT_DPDO(t *testing.T) {
	d := NewDistBT(4)

	fmt.Println("base", d.base)
	mem := make([][]int, d.n, d.n)
	for i := 0; i < d.n; i++ {
		mem[i] = make([]int, d.n, d.n)
		for j := 0; j < d.n; j++ {
			mem[i][j] = -1
		}
	}
	d.mem = mem
	d.minDist = d.dpdo(d.n-1, d.n-1)
	fmt.Println(d.minDist)
	fmt.Println(d.mem)
}

func TestMaxIncSubf(t *testing.T) {
	a := "2936517"
	MaxIncSubf([]byte(a))
}

func Test_A(t *testing.T) {
	a := "    D,   A"
	fmt.Println(isPalindrome(a))
}

func isPalindrome(s string) bool {
	b := []byte(s)
	slen := len(s)
	head, end := 0, slen-1
	for head < end {
		fmt.Println("head, end", head, end, b)
		if 'A' <= b[head] && b[head] <= 'Z' {
			b[head] += 'a' - 'A'
		}
		if 'A' <= b[end] && b[end] <= 'Z' {
			b[end] += 'a' - 'A'
		}
		fmt.Println(!((b[head] >= '0' && b[head] <= '9') || (b[head] >= 'a' && b[head] <= 'z')))
		if !((b[head] >= '0' && b[head] <= '9') || (b[head] >= 'a' && b[head] <= 'z')) {
			head++
			continue
		}
		if !((b[end] >= '0' && b[end] <= '9') || (b[end] >= 'a' && b[end] <= 'z')) {
			end--
			continue
		}
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
