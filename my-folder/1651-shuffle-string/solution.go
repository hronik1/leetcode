func restoreString(s string, indices []int) string {
    out := make([]byte, len(s))
    
    for i := 0; i < len(indices); i++ {
        out[indices[i]] = s[i]
    }
    
    return string(out)
}
