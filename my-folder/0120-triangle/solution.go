import "fmt"
func minimumTotal(triangle [][]int) int {
    memo := map[string]int{}
    return minimumTotalHelper(triangle, 0, 0, memo)
}

func minimumTotalHelper(triangle [][]int, row int, col int, memo map[string]int) int {
    key := fmt.Sprintf("%d_%d", row, col)
    if v, ok := memo[key]; ok {
        return v
    }

    out := triangle[row][col]
    if row == len(triangle)-1 {
        return out
    }

    minLen := minimumTotalHelper(triangle, row+1, col, memo)
    rightLen := minimumTotalHelper(triangle, row+1, col+1, memo)
    if rightLen < minLen {
        minLen = rightLen
    }

    out += minLen
    memo[key] = out

    return out
}
