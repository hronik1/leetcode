func maxAreaOfIsland(grid [][]int) int {
    maxArea := 0
    visited := make([][]bool, len(grid))
    for i := range visited {
        visited[i] = make([]bool, len(grid[0]))
    }
    
    for row := 0; row < len(grid); row++ {
        for col := 0; col < len(grid[0]); col++ {
            if area := islandArea(row, col, grid, visited); area > maxArea {
                maxArea = area
            }
        }
    }
    
    return maxArea
}

func islandArea(row int, col int, grid [][]int, visited[][]bool) int{
    // bounds checks
    if row < 0 || row >= len(grid) || col < 0 || col >= len(grid[0]) {
        return 0
    }
    
    if visited[row][col] {
        return 0
    }
    
    visited[row][col] = true
    
    if grid[row][col] == 0 {
        return 0
    }
    
    return 1 + islandArea(row-1, col, grid, visited) + islandArea(row+1, col, grid, visited) + islandArea(row, col-1, grid, visited) + islandArea(row, col+1, grid, visited)
}
