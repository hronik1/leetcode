func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    sCounts := byteCounts(s)
    tCounts := byteCounts(t)
    for b, sCount := range sCounts {
        if tCount, ok := tCounts[b]; ok {
            if tCount != sCount {
                return false
            }
        } else {
            return false
        }
    }

    return true
}

func byteCounts(s string) map[rune]int {
    mapping := map[rune]int{}
    for _, v := range s {
        if count, ok := mapping[v]; ok {
            mapping[v] = count+1
        } else {
            mapping[v] = 1
        }
    }

    return mapping
}
