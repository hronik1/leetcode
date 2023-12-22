func longestCommonSubsequence(text1 string, text2 string) int {
    memo := [][]int{}
    for i := 0; i < len(text1); i++ {
        memo = append(memo, make([]int, len(text2), len(text2)))
        for j := 0; j < len(text2); j++ {
            memo[i][j] = -1
        }
    }

    return longestCommonSubsequenceStartingAt(text1, text2, 0, 0, memo)
}

func longestCommonSubsequenceStartingAt(text1 string, text2 string, i int, j int, memo [][]int) int {
    if i >= len(text1) || j >= len(text2) {
        return 0
    }
    
    if memo[i][j] >= 0 {
        return memo[i][j]
    }

    subsequenceLength := 0
    if text1[i] == text2[j] {
        subsequenceLength++
    }

    if i == len(text1)-1 && j == len(text2)-1 {
        return subsequenceLength
    }

    if text1[i] == text2[j] {
        subsequenceLength += longestCommonSubsequenceStartingAt(text1, text2, i+1, j+1, memo)
    } else {
        best := 0
        if i < len(text1)-1 {
            best = longestCommonSubsequenceStartingAt(text1, text2, i+1, j, memo)
        }

        other := 0
        if j < len(text2)-1 {
            other = longestCommonSubsequenceStartingAt(text1, text2, i, j+1, memo)
        }

        if other > best {
            best = other
        }

        subsequenceLength += best
    }

    memo[i][j] = subsequenceLength

    return subsequenceLength
}
