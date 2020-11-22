func spiralOrder(matrix [][]int) []int {
    n := len(matrix)
    if n == 0 {
        return []int{}
    }
    
    m := len(matrix[0])
    if m == 0 {
        return []int{}
    }
    
    stop := len(matrix)/2
    out := []int{}
    product := m*n
    count := 0
    for offset := 0; offset <= stop; offset++ {
        //TODO need validation on other?
        //add row going right
        for i, j := offset, offset; j < m-offset && count < product; j++ {
            out = append(out, matrix[i][j])
            count++
        }
        
        //add column going down
        for i, j := offset+1, m-1-offset; i < n-offset && count < product; i++ {
            out = append(out, matrix[i][j])
            count++
        }
        
        //add row going left
        for i, j := n-1-offset, m-2-offset; j >= offset && count < product; j-- {
            out = append(out, matrix[i][j])
            count++
        }
        
        //add column going down
        for i, j := n-2-offset, offset; i > offset && count < product; i-- {
            out = append(out, matrix[i][j])
            count++
        }
    }
    
    return out
}
