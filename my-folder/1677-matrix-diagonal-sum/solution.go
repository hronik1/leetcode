func diagonalSum(mat [][]int) int {
    l := len(mat)
    sum := 0
    for i := 0; i < l; i++ {
        sum += mat[i][i]
        sum += mat[i][l-1-i]
    }
    
    if l%2 == 1 {
        sum -= mat[l/2][l/2]
    }
    
    return sum
}
