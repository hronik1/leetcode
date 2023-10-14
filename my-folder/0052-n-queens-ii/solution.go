import "fmt"
import "math"
import "strconv"
import "strings"

type Cell struct {
    row int
    col int
}
func totalNQueens(n int) int {
    placedQueens := map[string]bool{}
    return totalNQueensHelper(n, 0, placedQueens)
}

func totalNQueensHelper(n int, row int, placedQueens map[string]bool) int {
    count := 0
    for col := 0; col < n; col += 1 {
        if canPlaceQueen(row, col, placedQueens) {
            if n == row+1 {
                return 1
            }
            
            placeQueen(row, col, placedQueens)
            count += totalNQueensHelper(n, row+1, placedQueens)
            removeQueen(row, col, placedQueens)
        }
    }
    
    return count
}

func canPlaceQueen(row int, col int, placedQueens map[string]bool) bool {
    for formattedQueen, _ := range placedQueens {
        placedCell := decodeCell(formattedQueen)
        if canAttack(row, col, placedCell.row, placedCell.col) {
            return false
        }
    }
    
    return true
}

func placeQueen(row int, col int, placedQueens map[string] bool) {
    formattedCell := formatCell(row, col)
    placedQueens[formattedCell] = true
}

func removeQueen(row int, col int, placedQueens map[string] bool) {
    formattedCell := formatCell(row, col)
    delete(placedQueens, formattedCell)
}

func formatCell(row int, col int) string {
    return fmt.Sprintf("%d_%d", row, col)
}

func decodeCell(formattedCell string) Cell {
    decodedCell := strings.Split(formattedCell, "_")
    row, _ := strconv.Atoi(decodedCell[0])
    col, _ := strconv.Atoi(decodedCell[1])
    return Cell{row:row, col:col}
}

func canAttack(row1 int, col1 int, row2 int, col2 int) bool {
    return row1 == row2 || col1 == col2 || math.Abs(float64(row1-row2)) == math.Abs(float64(col1-col2))
}
