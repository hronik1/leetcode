func lengthOfLongestSubstring(s string) int {
    longestLength := 0
    currentLength := 0
    seen := map[string]int{}
    substringStartI := 0
    
    for i, c := range s {
        if seenI, ok := seen[string(c)]; ok && seenI >= substringStartI {
            if currentLength > longestLength {
                longestLength = currentLength
            }
            currentLength = currentLength - (seenI + 1 - substringStartI)
            substringStartI = seenI+1
        }
        
        currentLength++
        seen[string(c)] = i
    }
    
    if currentLength > longestLength {
        longestLength = currentLength
    }
    
    return longestLength
}
