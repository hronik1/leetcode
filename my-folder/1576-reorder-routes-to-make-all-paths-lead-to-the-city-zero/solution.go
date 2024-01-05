import "fmt"
func minReorder(n int, connections [][]int) int {
    forwardNeighbors := make([][]int, n, n)
    reverseNeighbors := make([][]int, n, n)
    for _, connection := range connections {
        forwardNeighbors[connection[0]] = append(forwardNeighbors[connection[0]], connection[1])
        reverseNeighbors[connection[1]] = append(reverseNeighbors[connection[1]], connection[0])
    }

    visited := make([]bool, n, n)
    visited[0] = true
    
    curLevel := []int{0}
    changedCount := 0
    for len(curLevel) > 0 {
        nextLevel := []int{}
        for _, city := range curLevel {
            for _, forwardNeighbor := range forwardNeighbors[city] {
                if !visited[forwardNeighbor] {
                    nextLevel = append(nextLevel, forwardNeighbor)
                    visited[forwardNeighbor] = true
                    changedCount++
                }
            }

            for _, reverseNeighbor := range reverseNeighbors[city] {
                if !visited[reverseNeighbor] {
                    nextLevel = append(nextLevel, reverseNeighbor)
                    visited[reverseNeighbor] = true
                }
            }
        }
        curLevel = nextLevel
    }

    return changedCount
}
