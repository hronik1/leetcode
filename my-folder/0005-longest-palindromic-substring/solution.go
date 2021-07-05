func longestPalindrome(s string) string {
    if s == "" {
        return ""
    }
    
    start, end := 0, 0
    for i := 0; i < len(s); i++ {
        l, h := expandAroundCenter(s, i, i)
        if h-l > end-start {
            start, end = l, h
        }
        
        l, h = expandAroundCenter(s, i, i+1)
        if h-l > end-start {
            start, end = l, h
        }
    }
    
    return s[start: end+1]
}

// lo, hi inclusive
func expandAroundCenter(s string, lo int, hi int) (int, int) {
    if hi >= len(s) {
        return lo, lo
    }
    
    l, h := lo, hi
    for l >= 0 && h < len(s) && s[l] == s[h] {
        l--
        h++
    }
    
    return l+1, h-1
}
