func isSubsequence(s string, t string) bool {
    if len(s) > len(t) {
        return false
    }

    i := 0
    for j := 0; i < len(s) && j < len(t); j += 1 {
        if s[i] == t[j] {
            i += 1
        }
    }

    if i < len(s) {
        return false
    }
 
    return true
}
