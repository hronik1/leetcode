func minPartitions(n string) int {
    out := 0
    for _, r := range n {
        if diff := int(r - '0'); diff > out {
            out = diff
        } 
    }
    
    return out
}
