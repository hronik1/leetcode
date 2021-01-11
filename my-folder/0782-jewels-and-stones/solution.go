func numJewelsInStones(jewels string, stones string) int {
    j := map[rune]bool{}
    for _, r := range jewels {
        j[r] = true
    }
    
    out := 0
    for _, r := range stones {
        if ok := j[r]; ok {
            out++
        }
    }
    
    return out
}
