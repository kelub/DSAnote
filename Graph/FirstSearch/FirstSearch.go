package FirstSearch

import(
	"kelub/DSAnote/Queue"
	//"fmt"
	//"container/list"
	"fmt"
)

// 无向图
type Graph struct {
	vertex 	int 		//顶点数
	adj 	[][]int	//邻接表
}

func New(v int) *Graph{
	adj := make([][]int,v,v)
	for i := 0;i<v;i++{
		adj[i] = make([]int,0)
	}
	return &Graph{
		vertex: v,
		adj: adj,
	}
}

func (g *Graph)addEdge(s ,t int){

	g.adj[s] = append(g.adj[s],t)
	g.adj[t] = append(g.adj[t],s)
}

/*
0

1   2   4

3   5   6

    7
*/
func bulidTest() *Graph{
	g := New(8)
	g.addEdge(0,1)
	g.addEdge(1,2)
	g.addEdge(1,3)
	g.addEdge(2,4)
	g.addEdge(2,5)

	g.addEdge(4,6)

	g.addEdge(3,5)
	g.addEdge(5,6)
	g.addEdge(5,7)
	return g
}

//⼴度优先搜索（Breadth-First-Search）
func (g *Graph)BFS(s,t int) bool {
	if s == t{
		return true
	}
	// 已访问顶点记录
	visited := make([]bool,g.vertex,g.vertex)
	// 路径记录
	prev := make([]int,g.vertex,g.vertex)
	for i := range prev{
		prev[i] = -1
	}
	q := Queue.New(20)
	q.Enqueue(s)
	for q.Size() != 0{
		v := q.Dequeue()
		if v == nil{
			break
		}
		adj := g.adj[v.(int)]
		for i := 0;i<len(adj);i++{
			if !visited[adj[i]]{
				prev[adj[i]] = v.(int)
				if adj[i] == t{
					fmt.Println(prev)
					return true
				}
			}
			visited[adj[i]] = true
			q.Enqueue(adj[i])
		}
	}
	return true
}

// 深度优先搜索（Depth-First-Search）

func (g *Graph)DFS(s,t int){
	if s == t {
		return
	}
	visited := make([]bool,g.vertex,g.vertex)
	prev := make([]int,g.vertex,g.vertex)
	for i := range prev{
		prev[i] = -1
	}
	g.dfs(s,t,visited,prev)
	fmt.Println(prev)
}

func (g *Graph) dfs(s,t int,visited []bool,prev []int) bool{
	visited[s] = true
	if s == t {
		return true
	}
	adj := g.adj[s]
	for i := 0;i<len(adj);i++{
		if !visited[adj[i]]{
			prev[adj[i]] = s
			g.dfs(adj[i],t,visited,prev)
		}

	}
	return false
}