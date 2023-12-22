func uniquePaths(m int, n int) int {
    memo := [][]int{}
    for i := 0; i < m; i++ {
        memo = append(memo, make([]int, n, n))
    }

    return uniquePathsMemo(m, n, 0, 0, memo)
}

func uniquePathsMemo(m int, n int, i int, j int, memo [][]int) int {
    if i == m-1 && j == n-1 {
        return 1
    }

    if memod := memo[i][j]; memod > 0 {
        return memod
    }

    pathsFromHere := 0
    if i < m - 1 {
        pathsFromHere += uniquePathsMemo(m, n, i+1, j, memo)
    }

    if j < n - 1 {
        pathsFromHere += uniquePathsMemo(m, n, i, j+1, memo)
    }

    memo[i][j] = pathsFromHere
    return pathsFromHere
}

