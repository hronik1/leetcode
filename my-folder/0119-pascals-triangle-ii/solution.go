func getRow(rowIndex int) []int {
    if rowIndex == 0 {
        return []int{1}
    }
    
    if rowIndex == 1 {
        return []int{1,1}
    }
    
    prevRow := getRow(rowIndex-1)
    out := make([]int, rowIndex+1)
    out[0] = 1
    out[rowIndex] = 1
    for i := 1; i < rowIndex; i += 1 {
        out[i] = prevRow[i-1] + prevRow[i]
    }
    
    return out
}
