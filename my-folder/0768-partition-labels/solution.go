func partitionLabels(S string) []int {
    lastSeen := map[rune]int{}
    for i, r := range S {
        lastSeen[r] = i
    }
    
    partLo, partHi := 0, 0
    out := []int{}
    for i, r := range S {
        if lastSeen[r] > partHi {
            partHi = lastSeen[r]
        }
        
        if i == partHi {
            out = append(out, partHi-partLo+1)
            partLo, partHi = i+1, i+1
        }
    }
    
    return out
}
