func flipAndInvertImage(A [][]int) [][]int {
    for i := 0; i < len(A); i++ {
        l := len(A[i])
        stop := l/2
        for j := 0; j < stop; j++ {
            temp := A[i][j]
            A[i][j] = A[i][l-j-1]
            A[i][l-j-1] = temp
        }
    } 
    
    for i := 0; i < len(A); i++ {
        l := len(A[i])
        for j := 0; j < l; j++ {
            if A[i][j] == 0 {
                A[i][j] = 1
            } else {
                A[i][j] = 0
            }
        }
    }
    
    return A
}
