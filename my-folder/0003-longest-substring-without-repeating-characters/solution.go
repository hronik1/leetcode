func lengthOfLongestSubstring(s string) int {
    maxL := 0
    startI := 0
    seen := map[rune]int{}
    
    for i, r := range s {
        if j, ok := seen[r]; ok {
            if j >= startI {
                startI = j+1
            }
        } 
        seen[r] = i
        
        if i-startI+1 > maxL {
            maxL = i-startI+1
        }
    }
    
    return maxL
}
