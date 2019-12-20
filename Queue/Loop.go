package Queue

import (
	"sync/atomic"
	// "fmt"
)

/*
循环队列 重点判断队列为空或者满。
*/
type LoopQueue struct {
	cap int64

	items []interface{}

	tail int64
	head int64
}

func New(cap int64) *LoopQueue {
	return &LoopQueue{
		cap:   cap,
		items: make([]interface{}, cap, cap),
		tail:  0,
		head:  0,
	}
}

// Enqueue enter queue
func (q *LoopQueue) Enqueue(item interface{}) bool {
	var tail int64
	tail = atomic.LoadInt64(&q.tail)
	if (tail+1)%q.cap == atomic.LoadInt64(&q.head) {
		return false
	}
	for {
		tail = atomic.LoadInt64(&q.tail)
		if atomic.CompareAndSwapInt64(&q.tail, tail, (tail+1)%q.cap) {
			q.items[tail] = item
			break
		}
	}
	// atomic.CompareAndSwapInt64(&q.tail,tail,(tail + 1) % q.cap)
	// q.tail = (q.tail + 1) % q.cap
	return true
}

// Dequeue de queue
func (q *LoopQueue) Dequeue() interface{} {
	// fmt.Println("head tail",q.head,q.tail)
	var head int64
	head = atomic.LoadInt64(&q.head)
	if head == atomic.LoadInt64(&q.tail) {
		return nil
	}
	for {
		head = atomic.LoadInt64(&q.head)
		if atomic.CompareAndSwapInt64(&q.head, head, (head+1)%q.cap) {
			return q.items[head]
		}
	}
	// r := q.items[head]
	// atomic.CompareAndSwapInt64(&q.head,head,(head + 1) % q.cap)
	// return r
}

// Size return queue use size
func (q *LoopQueue) Size() int64 {
	r := atomic.LoadInt64(&q.tail) - atomic.LoadInt64(&q.head)
	if r >= 0 {
		return r
	}
	return r + q.cap + 1

}
