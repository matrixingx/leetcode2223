package test5

import "fmt"

    
func main() {
    var n,m = 3,4
    var matrix = make([][]int,n)
    for i := range matrix {
        matrix[i] = make([]int, m)
    }
    var dir = [][]int{{-1,1},{1,-1}}
    var wall = initWall(n,m)
    var dirIndex = 1
    var i,j = 0,0
    for count := 1 ; count <= n*m ; count ++ {
        matrix[i][j] = count
        hitWall,wallType := isWall(wall,i,j)
        if hitWall {
            if wallType == 3 {
                i,j = i+1,j
            }
            if wallType == 2 {
                i,j = i,j+1
            }
            dirIndex = (dirIndex+1)%2
        } else {
            i,j = i+dir[dirIndex][0],j+dir[dirIndex][1]
        }
    }
    fmt.Println(matrix)
}

func initWall(n,m int) [][]int {
    var wall = make([][]int,3)
    for i := range wall {
        wall[i] = make([]int, 4)
    }
    for j := 1 ; j < m-1 ; j=j+2 {
        wall[0][j] = 2 // 用2来代表上下边界，往右移动
    }
    for j := 0 ; j < m ; j=j+2 {
        wall[n-1][j] = 2 // 用2来代表上下边界，往右移动
    }
    for i := 0 ; i < n-1 ; i=i+2 {
        wall[i][0] = 3 // 用3来代表左右边界，往下移动
    }
	for i := 0 ; i < n ; i=i+2 {
		wall[i][m-1] = 3
	}
    return wall
}

func isWall(wall [][]int,i,j int) (bool,int) {
    if wall[i][j] != 0 {
        return true,wall[i][j]
    }
    return false,0
}