func xorOperation(n int, start int) int {
    out := 0
    for i, cur := 0, start; i < n; i, cur = i+1, cur+2 {
        out ^= cur
    }
    
    return out
}
