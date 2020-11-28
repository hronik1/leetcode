func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    
    sRuneCount := map[rune]int{}
    for _, r := range s {
        sRuneCount[r]++
    }
    
    tRuneCount := map[rune]int{}
    for _, r := range t {
        tRuneCount[r]++
    }
    
    for r, sCount := range sRuneCount {
        if tCount, ok := tRuneCount[r]; ok {
            if tCount != sCount {
                return false
            }
        } else {
            return false
        }
    }
    
    return true
}
