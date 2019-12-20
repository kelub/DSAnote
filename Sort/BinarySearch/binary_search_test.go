package BinarySearch

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func Test_BSSliceSqrt(t *testing.T) {
	bbs := &BSSlice{}
	num := float64(10)
	r, err := bbs.Sqrt(num)
	assert.Nil(t, err)
	assert.True(t, math.Abs(r-math.Sqrt(num)) < 0.000001)
}

func Test_BSSliceFirstValue(t *testing.T) {
	bbs := &BSSlice{}
	var actual [][]int
	var expected []int
	var value int = 6

	actual = [][]int{[]int{1, 5, 5, 7, 7, 7, 7, 9}, []int{1, 2, 2, 2, 2, 2, 2, 6}, []int{6, 6, 6, 6, 6, 7, 8, 9}, []int{6, 6}}
	expected = []int{-1, 7, 0, 0}

	for i := 0; i < len(actual); i++ {
		index := bbs.FirstValue(actual[i], value)
		assert.Equal(t, index, expected[i])
	}

}
