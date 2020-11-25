func minDistance(word1 string, word2 string) int {
    memo := make(map[int]map[int]int)
    for i := 0; i <= len(word1); i++ {
        memo[i] = make(map[int]int)
    }
    
    fmt.Printf("%v", memo)
    return minDistanceHelper(word1, word2, 0, 0, memo)
}

func minDistanceHelper(word1 string, word2 string, i int, j int, memo map[int]map[int]int) int {
    if outer, outerOk := memo[i]; outerOk {
        if res, ok := outer[j]; ok {
            return res
        }
    }
    
    if i == len(word1) {
        out := len(word2)-j
        memo[i][j] = out
        return out
    }
    
    if j == len(word2) {
        out := len(word1)-i
        memo[i][j] = out
        return out
    }
    
    if word1[i] == word2[j] {
        out := minDistanceHelper(word1, word2, i+1, j+1, memo)
        memo[i][j] = out
        return out
    }
    
    out := minDistanceHelper(word1, word2, i+1, j+1, memo)
    if out1 := minDistanceHelper(word1, word2, i+1, j, memo); out1 < out {
        out = out1
    }
    if out1 := minDistanceHelper(word1, word2, i, j+1, memo); out1 < out {
        out = out1
    }
    out++
    
    memo[i][j] = out
    return out
}
