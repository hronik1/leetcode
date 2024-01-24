func finalValueAfterOperations(operations []string) int {
    out := 0
    for _, s := range operations {
        if s[0] == '-' || s[2] == '-' {
            out--
        } else {
            out++
        }
    }

    return out
}
