func setZeroes(matrix [][]int)  {
    zeroRows := map[int]bool{}
    zeroColumns := map[int]bool{}
    for i, row := range matrix {
        for j, v := range row {
            if v == 0 {
                zeroRows[i] = true
                zeroColumns[j] = true
            }
        }
    }

    for i := 0; i < len(matrix); i += 1 {
        if _, ok := zeroRows[i]; ok {
            for j := 0; j < len(matrix[i]); j += 1 {
                matrix[i][j] = 0
            }
        }
    }

    for j := 0; j < len(matrix[0]); j += 1 {
        if _, ok := zeroColumns[j]; ok {
            for i := 0; i < len(matrix); i += 1 {
                matrix[i][j] = 0
            }
        }
    }
}
