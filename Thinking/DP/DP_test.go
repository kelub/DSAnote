package DP

import (
	"fmt"
	"testing"
)

func Test_do(t *testing.T) {
	p := &BackPack{
		items:  []int{1, 2, 4, 8, 16, 32},
		weight: 100,
	}
	p.do(0, 0, p.weight, len(p.items), p.items)
	fmt.Println(p.maxWeight)
}

func Test_domem(t *testing.T) {
	p := &BackPack{
		items:  []int{1, 2, 4, 8, 16, 32},
		weight: 100,
	}
	p.domem(0, 0, p.weight, len(p.items), p.items)
	fmt.Println(p.maxWeight)
}

func Test_dp(t *testing.T) {
	p := &BackPack{
		items:  []int{2, 2, 4, 6, 3},
		weight: 9,
	}
	p.dp(p.weight, len(p.items))
	fmt.Println(p.maxWeight)
}

func Test_dpValue(t *testing.T) {
	p := &BackPack{
		items:  []int{2, 2, 4, 6, 3},
		weight: 9,

		values: []int{3, 4, 8, 9, 6},
	}
	p.dpValue(p.weight, len(p.items))
	fmt.Println(p.maxValues)
}

func Test_YH(t *testing.T) {
	y := NewYH()
	fmt.Println("y", y)
	y.back(0, y.head, 0, y.layer)
	fmt.Println("minPath", y.minPath)
}

func Test_YHDP(t *testing.T) {
	y := NewYH()
	fmt.Println("y", y)
	y.dp(y.head, y.layer)
	fmt.Println("minPath", y.minPath)
}
