func lengthOfLongestSubstring(s string) int {
    seen := map[byte]bool{}
    bestLen := 0
    curLen := 0
    for i := 0; i < len(s); i += 1 {
        if _, ok := seen[s[i]]; ok {
            for j := i - curLen; j < i; j += 1 {
                delete(seen, s[j])
                curLen -= 1
                if s[i] == s[j] {
                    break
                }
            }
        }

        seen[s[i]] = true
        curLen += 1
        if curLen > bestLen {
            bestLen = curLen
        }
    }

    return bestLen
}
