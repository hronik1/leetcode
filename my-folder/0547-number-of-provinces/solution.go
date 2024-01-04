func findCircleNum(isConnected [][]int) int {
    parents := make([]int, len(isConnected), len(isConnected))
    for i := 0; i < len(parents); i++ {
        parents[i] = i
    }

    //fmt.Printf("parents: %d\n", parents)
    for i := 0; i < len(isConnected)-1; i++ {
        for j := i+1; j < len(isConnected); j++ {
            if isConnected[i][j] == 1 {
                union(parents, i, j)
            }
        }
    }
 

    //fmt.Printf("parents: %d\n", parents)
    parentsSeen := map[int]bool{}
    for _, parent := range parents {
        f := find(parents, parent)
        if _, ok := parentsSeen[f]; !ok {
            parentsSeen[f] = ok
        }
    }

    return len(parentsSeen)
}

func find(parents[]int, i int) int {
    if parents[i] == i {
        return i
    } else {
        res := find(parents, parents[i])
        parents[i] = res
        return res
    }
}

func union(parents[]int, i int, j int) {
    parentI := find(parents, i)
    parentJ := find(parents, j)
    parents[parentI] = parentJ
}
