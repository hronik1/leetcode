func decode(encoded []int, first int) []int {
    out := make([]int, len(encoded)+1)
    out[0] = first
    for i := 0; i < len(encoded); i++ {
        out[i+1] = out[i] ^ encoded[i]
    }
    
    return out
}
