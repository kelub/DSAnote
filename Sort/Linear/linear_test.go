package Linear

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_CountingSort(t *testing.T) {
	var actual [][]int
	var expected [][]int

	actual = [][]int{[]int{3, 6, 2, 2, 1, 4, 0, 5}, []int{3, 8, 2, 2, 9, 4, 0, 5}, []int{3, 6, 2, 2, 0, 4, 0, 5, 0, 0, 0, 0, 0}}
	expected = [][]int{[]int{0, 1, 2, 2, 3, 4, 5, 6}, []int{0, 2, 2, 3, 4, 5, 8, 9}, []int{0, 0, 0, 0, 0, 0, 0, 2, 2, 3, 4, 5, 6}}
	cs := New()

	for i := 0; i < len(actual); i++ {
		cs.CountingSort(actual[i])
		assert.Equal(t, actual[i], expected[i])
	}

}
