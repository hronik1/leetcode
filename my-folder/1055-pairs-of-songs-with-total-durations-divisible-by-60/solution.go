func numPairsDivisibleBy60(time []int) int {
    counts := make([]int, 60)
    out := 0
    for _, t := range time {
        m := t%60
        complement := (60-m)%60
        out += counts[complement]
        counts[m]++
    }
    
    return out
}
