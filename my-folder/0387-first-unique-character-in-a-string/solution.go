func firstUniqChar(s string) int {
    charCount := map[rune]int{}
    for _, v := range s {
        charCount[v] += 1 
    }
    
    for i, v := range s {
        if count, ok := charCount[v]; ok {
            if count == 1 {
                return i
            }  
        } 
    }
    
    return -1
}
