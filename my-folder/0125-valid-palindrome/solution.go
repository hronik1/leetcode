func isPalindrome(s string) bool {
    lo, hi := 0, len(s)-1
    for lo < hi {
        first := rune(s[lo])
        if !(first >= 'a' && first <= 'z') && !(first >= 'A' && first <= 'Z') && !(first >= '0' && first <= '9') {
            lo++
            continue
        }
        
        last := rune(s[hi])
        if !(last >= 'a' && last <= 'z') && !(last >= 'A' && last <= 'Z') && !(last >= '0' && last <= '9') {
            hi--
            continue
        }

        if first >= 'A' && first <= 'Z' {
            first = 'a' + (first-'A')
        }
        if last >= 'A' && last <= 'Z' {
            last = 'a' + (last-'A')
        }
        
        if first != last {
            return false
        }
        
        lo++
        hi--
    }
    
    return true
}
