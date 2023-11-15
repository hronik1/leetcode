func longestCommonPrefix(strs []string) string {
    for i, v := range strs[0] {
        for _, str := range strs {
            if i >= len(str) || str[i] != byte(v) {
                return strs[0][0:i]
            } 
        }

        
    }
    return strs[0]
}
