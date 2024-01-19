func partitionString(s string) int {
    if len(s) == 0 {
        return 0
    }

    partitions := 1
    seen := map[rune]bool{}
    for _, r := range s {
        if _, ok := seen[r]; ok {
            partitions++
            seen = map[rune]bool{}
        }

        seen[r] = true
    }

    return partitions
}
