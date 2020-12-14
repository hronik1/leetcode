func partition(s string) [][]string {
    out := [][]string{}
    for i := 1; i <= len(s); i++ {
        ss := s[0:i]
        if isPalindrome(ss) {
            subsolutions := partition(s[i:])
            if len(subsolutions) == 0 {
                out = append(out, []string{ss})
            } else {
                for _, sub := range subsolutions {
                    appended := append([]string{ss}, sub...)
                    out = append(out, appended)
                }
            }
        }
    }
    
    return out
}

func isPalindrome(s string) bool {
    for i := 0; i < len(s); i++ {
        if s[i] != s[len(s)-1-i] {
            return false
        }
    }
    
    return true
}
