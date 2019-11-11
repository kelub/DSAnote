package Queue

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Enqueue(t *testing.T) {
	q := New(7)

	r := q.Enqueue("a")
	assert.True(t, r)
	r = q.Enqueue("b")
	assert.True(t, r)
	de := q.Dequeue()
	assert.NotNil(t, de)
	de = de.(string)
	assert.Equal(t, de, "a")
	q.Dequeue()
	assert.Equal(t, q.Size(), int64(0))
	de = q.Dequeue()
	assert.Nil(t, de)

	q.Enqueue("a")
	q.Enqueue("b")
	q.Enqueue("c")
	q.Enqueue("d")
	q.Enqueue("e")
	q.Enqueue("f")
	q.Enqueue("g")

	assert.False(t, q.Enqueue("h"))
	q.Dequeue()
	q.Dequeue()
	q.Dequeue()
	q.Enqueue("h")
	q.Enqueue("j")
	assert.Equal(t, q.Size(), int64(6))
}
