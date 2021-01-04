func minSteps(s string, t string) int {
    sCounts := map[rune]int {}
    for _, r := range s {
        val := 1
        if count, ok := sCounts[r]; ok {
            val += count
        }
        
        sCounts[r] = val
    }
    
    tCounts := map[rune]int {}
    for _, r := range t {
        val := 1
        if count, ok := tCounts[r]; ok {
            val += count
        }
        
        tCounts[r] = val
    }
    
    diff := 0
    for r, sCount := range sCounts {
        tCount := tCounts[r]
        diff += int(math.Abs(float64(sCount-tCount)))
    }
    
    for r, tCount := range tCounts {
        sCount := sCounts[r]
        // only count missing entries, so we don't double count
        if sCount == 0 {
           diff += tCount 
        }
    }
    
    return diff/2
}
