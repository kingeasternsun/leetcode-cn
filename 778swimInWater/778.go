package leetcode

/*
1.  因为可以在泳池中瞬间移动，那么就可以等到特定的时间后再移动
2.  如果在时间t的时候存在一条路游到终点，那么就可以从起点瞬间移动到终点
*/
func swimInWater(grid [][]int) int {

	g := Grid{G: grid, N: len(grid)}
	if g.N == 0 {
		return 0
	}

	start := grid[0][0]
	end := g.N * g.N
	for start < end {
		mid := (end-start)/2 + start
		visited := make([]bool, g.N*g.N)
		if g.DFS(0, 0, mid, visited) {
			end = mid
		} else {
			start = mid + 1
		}

	}

	return start

}

type Grid struct {
	G [][]int
	N int
}

var dir = []int{0, 1, 0, -1, 0}

//DFS 判断i时间是否可以
func (g *Grid) DFS(x, y int, t int, visited []bool) bool {
	if x == g.N-1 && y == g.N-1 {
		return true
	}

	for i := 0; i < 4; i++ {
		nx := x + dir[i]
		ny := y + dir[i+1]

		if nx < 0 || nx >= g.N || ny < 0 || ny >= g.N {
			continue
		}

		if visited[nx*g.N+ny] {
			continue
		}

		if g.G[nx][ny] <= t {
			visited[nx*g.N+ny] = true
			if g.DFS(nx, ny, t, visited) {
				return true
			}
		}

	}

	return false
}



import (
    "fmt"
)
func main() {
    //a := 0
    //fmt.Scan(&a)
    //fmt.Printf("%d\n", a)
    fmt.Printf("Hello World!\n");
    fmt.Println(judgeSub("123123",2))
    fmt.Println(judge("123123"))
}

func judge(s string)bool{
    
    sLen:=len(s)
    if sLen == 1{
        return true 
    }
    
 
    for i:=1;i<=sLen/2;i++{
        
        if sLen%i>0{
            continue
        }
        
        if judgeSub(s,i){
            return true
        }
    }
    
    return false
}

func judgeSub(s string, subLen int)bool{
    m:=len(s)/subLen
    
        for j:=0;j<subLen;j++{
            a:=s[j]
            for i:=0;i<m;i++{
                if s[i*subLen+j]!=a{
                    return false
                }
            }
        }
    
    return true
    
}