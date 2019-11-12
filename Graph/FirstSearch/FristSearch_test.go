package FirstSearch

import(
	"testing"
	"fmt"
)
func TestGraphBFS(t *testing.T){
	g := bulidTest()
	fmt.Println(g)
	for i := range g.adj{
		fmt.Printf("%d %v \n",i,g.adj[i])
	}
	g.BFS(0,7)
}

func TestGraphDFS(t *testing.T){
	g := bulidTest()
	fmt.Println(g)
	for i := range g.adj{
		fmt.Printf("%d %v \n",i,g.adj[i])
	}
	g.DFS(0,7)
}
