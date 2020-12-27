func maxIncreaseKeepingSkyline(grid [][]int) int {
    maxRows := make([]int, len(grid))
    maxColumns := make([]int, len(grid[0]))
    for i, row := range grid {
        for j, height := range row {
            if height > maxRows[i] {
                maxRows[i] = height
            }
            
            if height > maxColumns[j] {
                maxColumns[j] = height
            }
        }
    }
    
    diff := 0
    for i, row := range grid {
        for j, height := range row {
            newHeight := maxRows[i]
            if maxColumns[j] < newHeight {
                newHeight = maxColumns[j]
            }
            
            diff += (newHeight - height)
        }
    }
    
    return diff
}
