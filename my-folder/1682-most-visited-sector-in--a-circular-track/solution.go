func mostVisited(n int, rounds []int) []int {
    visited := []int{}
    if rounds[0] <= rounds[len(rounds)-1] {
        for i := rounds[0]; i <= rounds[len(rounds)-1]; i++ {
            visited = append(visited, i)
        }
    } else {
        for i := 1; i <= rounds[len(rounds)-1]; i++ {
            visited = append(visited, i)
        }
        for i := rounds[0]; i <= n; i++ {
            visited = append(visited, i)
        }
    }
    
    return visited
}
