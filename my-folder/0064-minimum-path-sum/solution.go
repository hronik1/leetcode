import "fmt"

func minPathSum(grid [][]int) int {
    memo := map[string]int{}
    return minPathSumHelper(grid, 0, 0, memo)
}

func minPathSumHelper(grid [][]int, row int, col int, memo map[string]int) int {
    key := fmt.Sprintf("%d_%d", row, col)
    if v, ok := memo[key]; ok {
        return v
    }

    out := grid[row][col]
    if row == len(grid)-1 && col == len(grid[row])-1 {
        memo[key] = out
        return out
    }

    downChildLen := -1
    if row < len(grid) - 1 {
        downChildLen = minPathSumHelper(grid, row+1, col, memo) 
    }

    rightChildLen := -1
    if col < len(grid[row]) - 1 {
        rightChildLen = minPathSumHelper(grid, row, col+1, memo)
    }

    minChildLen := downChildLen
    if (rightChildLen >= 0 && rightChildLen < minChildLen) || minChildLen < 0 {
        minChildLen = rightChildLen
    }

    out += minChildLen
    memo[key] = out

    return out
}
