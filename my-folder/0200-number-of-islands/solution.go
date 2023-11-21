func numIslands(grid [][]byte) int {
    if len(grid) == 0 {
        return 0
    }

    rows := len(grid) 
    columns := len(grid[0])
    visited := make([][]bool, rows)
    islands := 0
    for i := range visited {
        visited[i] = make([]bool, columns)
    }

    for r := 0; r < rows; r += 1 {
        for c := 0; c < columns; c += 1 {
            if !visited[r][c] && grid[r][c] == '1' {
                visited[r][c] = true
                visitIsland(grid, visited, r, c)
                islands += 1
            }
        }
    } 

    return islands
}

func visitIsland(grid [][]byte, visited [][]bool, r int, c int) {
    island := [][]int{{r, c}}
    directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
    for len(island) != 0 {
        cell := island[0]
        island = island[1:]
        for _, direction := range directions {
            row := cell[0] + direction[0]
            col := cell[1] + direction[1]
            if isInBounds(grid, row, col) && !visited[row][col] && grid[row][col] == '1' {
                visited[row][col] = true
                island = append(island, []int{row, col})
            }
        }
    }
}

func isInBounds(grid [][]byte, row int, col int) bool {
    return row >= 0 && row < len(grid) && col >= 0 && col < len(grid[0])
}
