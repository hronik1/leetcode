func allPathsSourceTarget(graph [][]int) [][]int {
    return listAllPaths(graph, 0, len(graph)-1)
}

func listAllPaths(graph [][]int, source int, target int) [][]int {
    if source == target {
        return [][]int{{target}}
    }
    
    out := [][]int{}
    for _, neighbor := range graph[source] {
        paths := listAllPaths(graph, neighbor, target)
        for _, path := range paths {
            path = append([]int{source}, path...)
            out = append(out, path)
        }
    }
    
    return out
}
