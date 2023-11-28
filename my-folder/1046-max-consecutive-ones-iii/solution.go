func longestOnes(nums []int, k int) int {
    startI := 0 // init this 
    bestStreak := 0
    zeroIndices := []int{}
    for i, v := range nums {
        if v == 0 {
            zeroIndices = append(zeroIndices, i)
            if len(zeroIndices) > k {
                startI = zeroIndices[0] + 1
                zeroIndices = zeroIndices[1:]
            }
        }

        curStreak := i + 1 - startI
        if curStreak > bestStreak {
            bestStreak = curStreak
        }
    }

    return bestStreak
}
