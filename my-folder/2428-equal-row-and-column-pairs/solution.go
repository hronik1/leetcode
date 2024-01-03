import "fmt"
import "strings"

func equalPairs(grid [][]int) int {
    rowsMap := buildRowsMap(grid)
    colsMap := buildColsMap(grid)

    matchingCount := 0
    for row, count := range rowsMap {
        if colCount, ok := colsMap[row]; ok {
            matchingCount += count*colCount
        }
    }

    return matchingCount
}

func buildRowsMap(grid [][]int) map[string]int {
    out := map[string]int {}
    for _, row := range grid {
        var b strings.Builder
        for _, cell := range row {
            fmt.Fprintf(&b, "%d_", cell)
        }
        s := b.String()
        out[s]++
    }

    return out
}

func buildColsMap(grid [][]int) map[string]int {
    out := map[string]int {}
    for j := 0; j < len(grid); j++ {
        var b strings.Builder
        for i := 0; i < len(grid); i++ {
            fmt.Fprintf(&b, "%d_", grid[i][j])
        }
        s := b.String()
        out[s]++
    }     

    return out
}
