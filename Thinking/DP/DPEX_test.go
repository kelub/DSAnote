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
