func isValidSudoku(board [][]byte) bool {
    // is each row valid
    for rowI := 0; rowI < len(board); rowI+=1 {
        seen := map[byte]struct{}{}
        for _, cell := range board[rowI] {
            if cell != '.' {
                if _, ok := seen[cell]; ok {
                    return false
                }
                
                seen[cell] = struct{}{}
            }
        }
    }
    
    // is each column valid
    for columnI := 0; columnI < len(board); columnI+=1 {
        seen := map[byte]struct{}{}
        for rowI := 0; rowI < len(board); rowI+=1 {
            cell := board[rowI][columnI]
            if cell != '.' {
                if _, ok := seen[cell]; ok {
                    return false
                }
                
                seen[cell] = struct{}{}
            }
        }
    }
    // is each cell valid
    for row := 0; row < 3; row+=1 {
        for column := 0; column < 3; column+=1 {
            if !isBoxValid(board, row, column) {
                return false
            }
        }
    }

    return true
}

func isBoxValid(board [][]byte, row int, column int) bool {
    seen := map[byte]struct{}{}
    for columnI := 3*column; columnI < 3*column+3; columnI+=1 {
        for rowI := 3*row; rowI < 3*row+3; rowI+=1 {
            cell := board[rowI][columnI]
            if cell != '.' {
                if _, ok := seen[cell]; ok {
                    return false
                }
                
                seen[cell] = struct{}{}
            }
        }
    }

    return true
}
