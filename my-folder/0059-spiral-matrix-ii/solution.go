func generateMatrix(n int) [][]int {
    out := make([][]int, n)
    for i := range out {
        out[i] = make([]int, n)
    }
    
    stop := n/2
    product := n*n
    count := 1
    
    for offset := 0; offset <= stop; offset++ {
        for i, j := offset, offset; count <= product && j < n-offset; j++ {
            out[i][j] = count
            count++
        }
        
        for i, j := offset+1, n-1-offset; count <= product && i < n-offset; i++ {
            out[i][j] = count
            count++
        }
        
        for i, j := n-1-offset, n-offset-2; count <= product && j >= offset; j-- {
            out[i][j] = count
            count++
        }
        
        for i, j := n-2-offset, offset; count <= product && i > offset; i-- {
            out[i][j] = count
            count++
        }
    }
    
    return out
}
