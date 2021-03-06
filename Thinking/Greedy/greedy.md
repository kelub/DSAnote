贪心算法，当前步最优解。一般不是全局最优解。
# 常见总结
定义一组数据，在有限制条件下，希望得到最优结果。同时一般满足，单步有最优解。
重点在于问题分析，找到一个方法满足，尽可能满足最小期望，也就是最小满足条件。

## 糖果问题：

m个糖果和n个孩⼦                    一组数据
每个孩子需求不同                    限制条件
需要尽可能满足更多的孩子（数量）     最优结果

分析
给需求最小的孩子分配恰好满足的糖果数量。
单步最好结果，每步结合也需满足最好结果。

## 区间问题
假设有一组区间                  一组数据
需要选出两两不相交的区间        限制条件
最多能选出多少个                最优结果
[6,8] [2,4] [3,5] [1,5] [5,9] [8,10]
不相交区间
[2,4] [6,8] [8,10]

分析
每次选择就尽量让右边区间尽量大。
单步最好结果，每步结合也需满足最好结果。

实例化问题
转化成整个大区间，每次选择都需让右区间最大。也就是左边区间每步占用最小。
每步处理一致，采用递归方式更加直观简单。
        1 2 3 4 5 6 7 8 9 10
        1       5
          2   4
            3   5
                5       9
                  6   8
                      8   10