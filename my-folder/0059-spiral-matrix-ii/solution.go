func generateMatrix(n int) [][]int {
    count := 1
    out := make([][]int, n)
    for i := 0; i < n; i++ {
        out[i] = make([]int, n)
    }
    
    for offset := 0; offset <= n/2; offset++ {
        //print top
        for i, j := offset, offset; j < n-offset; j++ {
            out[i][j] = count
            count++
        }
        
        //print right
        for i, j := offset+1, n-1-offset; i < n-offset; i++ {
            out[i][j] = count
            count++
        }
        
        //print bottom
        for i, j := n-1-offset, n-2-offset; j >= offset; j-- {
            out[i][j] = count
            count++
        }
        
        //print left
        for i, j := n-2-offset, offset; i > offset; i-- {
            out[i][j] = count
            count++
        }
    }
    
    return out
}
