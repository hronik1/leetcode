func canReach(arr []int, start int) bool {
    memo := map[int]bool{}
    visited := map[int]bool{}
    return canReachMemo(arr, start, memo, visited)
}

func canReachMemo(arr []int, start int, memo map[int]bool, visited map[int]bool) bool {
    if v, ok := memo[start]; ok {
        return v
    }

    visited[start] = true
    if arr[start] == 0 {
        memo[start] = true
        return true
    }

    leftI, rightI := start - arr[start], start + arr[start]
    canReachLeft := false
    if leftI >= 0 && !visited[leftI] {
        canReachLeft = canReachMemo(arr, leftI, memo, visited)
    }

    canReachRight := false
    if rightI < len(arr) && !visited[rightI] {
        canReachRight = canReachMemo(arr, rightI, memo, visited)
    }

    out := canReachLeft || canReachRight
    memo[start] = out

    return out
}
