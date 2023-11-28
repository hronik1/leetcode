func longestSubarray(nums []int) int {
    bestStreak := 0
    startI := 0
    prevZeroI := -1 
    for i, v := range nums {
        if v == 0 {
            startI = prevZeroI + 1
            prevZeroI = i
        }

        curStreak := i - startI
        if curStreak > bestStreak {
            bestStreak = curStreak
        }
    }

    return bestStreak
}
