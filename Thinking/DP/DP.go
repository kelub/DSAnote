package DP

// 1,2,5,10,15,20,25,50
type BackPack struct {
	items  []int // 每个物品重量数组
	weight int   // 物品总总量

	maxWeight int // 最大容量

	mem [][]bool

	// dpValue
	values    []int // 每个物品的价值数组
	maxValues int   // 最大价值

}

// do 普通暴力搜索
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

// domem
// do 复杂度高重复计算多
// 记录重复迭代，提高效率
func (p *BackPack) domem(i, curw, weight, n int, items []int) {
	if curw >= weight || i == n { // 背包满或者物品装完 结束
		if p.maxWeight < curw {
			p.maxWeight = curw
		}
		return
	}
	// 重复状态直接返回
	if p.mem[i][curw] == true {
		return
	}
	// 记录状态
	p.mem[i][curw] = true

	// 不装
	p.domem(i+1, curw, weight, n, items)
	// 装
	if weight >= curw+items[i] {
		p.domem(i+1, curw+items[i], weight, n, items)
	}
}

// dp
// 动态规划
// 这里采用一个状态数组，代替了前文的记录状态的数组
// 不管是暴力搜索或者采用记录数组还是动态规划，甚至递归非递归，最核心点在于，
// 分支的选择，也就是状态迁移过程。 这里表现为每个物品都有装或者不装的状态。
// 暴力递归 采用当前重量加上添加的重量，通过递归栈记录每个状态。
// 动态规划 采用数组记录每个状态,即所有可能的值，并通过状态继续迭代，直到最优解出现

func (p *BackPack) dp(weight, n int) {
	// 状态数组
	states := make([][]bool, n, n)
	for i := 0; i < n; i++ {
		states[i] = make([]bool, weight+1, weight+1)
	}
	// 初始化第0个物品
	states[0][0] = true
	states[0][p.items[0]] = true
	// 每层状态
	for i := 1; i < n; i++ {
		// 不装
		// 通循环加判断定位已经到达的状态也就是到达的当前重量
		// 重叠子问题
		// 记录这个状态
		for j := 0; j < weight; j++ {
			// 根据上一层状态推出下一个可能的状态，并记录这个状态
			// 这里有两个可能的状态要么装要么不装。
			if states[i-1][j] == true {
				states[i][j] = true
			}
		}
		// 装
		for j := 0; j < weight-p.items[i]; j++ {
			if states[i-1][j] == true {
				states[i][j+p.items[i]] = true
			}
		}
	}
	for m := n - 1; m >= 0; m-- {
		for i := weight; i >= 0; i-- {
			if states[m][i] == true {
				p.maxWeight = i
				return
			}
		}
	}

}

// dpValue
// 当问题再添加一个限制条件，不仅仅是最大容量，还有最大价值。暴力回溯加备忘录方式就无法实现了，
// 因为限制条件变成了最大值，采用DP可以方便舍弃不符合最大值情况
func (p *BackPack) dpValue(weight, n int) {
	// 状态数组 记录每个状态最大值
	states := make([][]int, n, n)
	for i := 0; i < len(states); i++ {
		states[i] = make([]int, weight+1, weight+1)
		for j := range states[i] {
			states[i][j] = -1
		}
	}

	// 初始化第0个物品
	// 不装
	states[0][0] = 0
	// 装
	states[0][p.items[0]] = p.values[0]

	for i := 1; i < n; i++ {
		// 不装
		for j := 0; j < weight; j++ {
			if states[i-1][j] >= 0 {
				states[i][j] = states[i-1][j]
			}
		}
		// 装
		for j := 0; j < weight-p.items[i]; j++ {
			if states[i-1][j] >= 0 {
				// 判断筛选，最大值
				// 在重量相同情况下，需要选择最大价值
				// 只有大于最大价值才会更新最大价值
				value := states[i-1][j] + p.values[i]
				if value > states[i][j+p.items[i]] {
					states[i][j+p.items[i]] = value
				}
			}
		}
	}

	for i := weight; i >= 0; i-- {
		if states[n-1][i] > p.maxValues {
			p.maxValues = states[n-1][i]
		}
	}
}
