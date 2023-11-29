func uniqueOccurrences(arr []int) bool {
    intCounts := map[int]int{}
    for _, v := range arr {
        intCounts[v]++
    }

    uniqueCounts := map[int]bool{}
    for _, v := range intCounts {
        if _, ok := uniqueCounts[v]; ok {
            return false
        }

        uniqueCounts[v] = true
    }

    return true
}
