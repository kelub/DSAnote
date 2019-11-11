package Heap

import "fmt"

type Heap struct {
	base  []int //底层数据
	cap   int   //堆容量
	count int   //已存储数据数量
}

func NewHeap(capacity int) *Heap {
	return &Heap{
		base: make([]int, capacity+1, capacity+1),
		cap:  capacity,
	}
}

// inster 插入数据  动态移动满足堆结构
func (h *Heap) inster(data int) bool {
	if h.count >= h.cap {
		return false
	}
	// h.base[0]不存数据，简洁代码 [i/2]即为[i]父节点
	h.count++
	h.base[h.count] = data

	//	从下往上移动，堆化
	i := h.count
	for (i/2 > 0) && h.base[i] > h.base[i/2] {
		temp := h.base[i/2]
		h.base[i/2] = h.base[i]
		h.base[i] = temp
		i = i / 2
	}

	return true
}

// dalete 删除堆顶元素，并堆化
func (h *Heap) remove() bool {
	if h.count == 0 {
		return false
	}

	// 从上往下
	h.count--
	index := 1
	h.base[index] = h.base[h.count]
	h.heapify(index)
	return true
}

func (h *Heap) heapify(index int) {
	for {
		maxPos := index
		// fmt.Println(index)
		if (2*index <= h.cap) && (h.base[index] < h.base[2*index]) {
			maxPos = 2 * index
		}
		if (2*index+1 <= h.cap) && (h.base[maxPos] < h.base[(2*index)+1]) {
			maxPos = 2*index + 1
		}
		if index == maxPos {
			return
		}
		//swap
		temp := h.base[index]
		h.base[index] = h.base[maxPos]
		h.base[maxPos] = temp
		index = maxPos
	}
}

//原地创建堆结构
func (h *Heap) buildHeap() {
	fmt.Println("firs ", h.base)
	for i := h.cap / 2; i >= 1; i-- {
		h.heapify(i)
		fmt.Println("i", i, " ", h.base)
	}
}
