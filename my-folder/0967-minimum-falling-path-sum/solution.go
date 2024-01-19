func minFallingPathSum(matrix [][]int) int {
    memo := map[string]int{}
    min := minFallingPathSumHelper(matrix, 0, 0, memo)
    for i := 1; i < len(matrix[0]); i++ {
        sum := minFallingPathSumHelper(matrix, 0, i, memo)
        if sum < min {
            min = sum
        }
    }

    return min
}

func minFallingPathSumHelper(matrix [][]int, row int, col int, memo map[string]int) int {
    key := fmt.Sprintf("%d_%d", row, col)
    if v, ok := memo[key]; ok {
        return v
    }

    sum := matrix[row][col]
    if row == len(matrix)-1 {
        memo[key] = sum
        return sum
    }

    minChild := minFallingPathSumHelper(matrix, row+1, col, memo)
    if col > 0 {
        left := minFallingPathSumHelper(matrix, row+1, col-1, memo)
        if left < minChild {
            minChild = left
        }
    }

    if col < len(matrix[row])-1 {
        right := minFallingPathSumHelper(matrix, row+1, col+1, memo)
        if right < minChild {
            minChild = right
        }
    }

    sum += minChild
    memo[key] = sum
    
    return sum
}
