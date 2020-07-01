package main

// import "math"

// 可以理解为 判断任意两个点之间是否相连 
func checkIfPrerequisite(n int, prerequisites [][]int, queries [][]int) []bool {

	dis:=make([][]bool,n)
	for i:=0;i<n;i++{
		dis[i] = make([]bool,n)
		// for j:=0;j<n;j++{

		// 	// if i!=j{
		// 	// 	dis[i][j] = math.MaxInt32
		// 	// }
		// }
		
	}


	for _,pre:=range prerequisites{

		dis[pre[0]][pre[1]]=true
	}

	for i:=0;i<n;i++{

		for j:=0;j<n;j++{

			
			for k:=0;k<n;k++{
				//这个迭代的顺利很重要 
				if dis[j][i]&& dis[i][k]{
					dis[j][k] = true
				}
			}
		}
	}


	result:=make([]bool,len(queries))
	for i,query:=range queries{
		result[i] = dis[query[0]][query[1]]
	}
	
	return result

}