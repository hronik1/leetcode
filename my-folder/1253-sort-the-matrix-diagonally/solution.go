import "sort"

func diagonalSort(mat [][]int) [][]int {
    startI, startJ := len(mat)-1, 0
    out := make([][]int, len(mat))
    for i := 0; i < len(mat); i++ {
        out[i] = make([]int, len(mat[i]))
    }
    
    for startJ < len(mat[0]) {
        items := []int{}
        for i, j := startI, startJ; i < len(mat) && j < len(mat[0]);  i, j = i+1, j+1 {
            items = append(items, mat[i][j])
        }
        
        sort.Slice(items, func(i, j int) bool {
            return items[i] < items[j]
        })
        
        for offset := 0; offset < len(items); offset++ {
            out[startI+offset][startJ+offset] = items[offset]
        }
        
        if startI > 0 {
            startI--
        } else {
            startJ++
        }
    }
    
    return out
}
