import "fmt"
import "strconv"

type Cell struct {
    I int
    J int
}

func solveSudoku(board [][]byte)  {
    // figure out empty cells
    emptyCells := []Cell{}
    for i, row := range board {
        for j, v := range row {
            if v == '.' {
                emptyCells = append(emptyCells, Cell{I:i, J:j})
            }
        }
    }
    
    // error check here if unsolvable?
    _ = solveSudokuHelper(board, emptyCells)
}

func solveSudokuHelper(board [][]byte, emptyCells []Cell) bool {
    if len(emptyCells) == 0 {
        return true
    }
    
    cell := emptyCells[0]
    for v := 1; v < 10; v += 1 {
        formattedVal := []byte(strconv.Itoa(v))[0]
        if isLegalPlacement(board, cell, formattedVal) {
            board[cell.I][cell.J] = formattedVal
            if solveSudokuHelper(board, emptyCells[1:]) {
                return true
            }
            board[cell.I][cell.J] = '.'
        }
    }
    
    return false
}

func isLegalPlacement(board [][]byte, cell Cell, val byte) bool {
    // validate column
    for i := 0; i < 9; i += 1 {
        if board[i][cell.J] == val {
            return false
        }
    }
    
    // validate row
    for j := 0; j < 9; j += 1 {
        if board[cell.I][j] == val {
            return false
        }
    }
    
    // validate Box
    boxI := (cell.I/3) * 3
    boxJ := (cell.J/3) * 3
    for i := boxI; i < boxI + 3; i += 1 {
        for j := boxJ; j < boxJ + 3; j += 1 {
            if board[i][j] == val {
                return false
            }
        }
    }
    
    return true
}
