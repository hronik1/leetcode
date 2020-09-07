import "strings"

func wordPattern(pattern string, str string) bool {
    words := strings.Split(str, " ")
    if len(pattern) != len(words) {
        return false
    }
    
    mapping := map[string]string{}
    reverseMapping := map[string]string{}
    for i := 0; i < len(pattern); i++ {
        c := string(pattern[i])
        if originalWord, ok := mapping[c]; ok {
            if originalWord != words[i] {
                return false
            } 
        }
        
        if originalC, ok := reverseMapping[words[i]]; ok {
            if originalC != c {
                return false
            }
        } 
        
        mapping[c] = words[i]
        reverseMapping[words[i]] = c
    }
    
    return true
}
