func minSubArrayLen(target int, nums []int) int {
    lo := 1
    hi := len(nums)
    curBest := -1
    for lo <= hi {
        windowSize := lo + (hi-lo)/2
        sum := maxWindowSizeSum(nums, windowSize)
        if sum >= target {
            curBest = windowSize
            hi = windowSize - 1
        } else {
            lo = windowSize + 1
        }

    }

    if curBest == -1 {
        return 0
    }

    return curBest
}

func maxWindowSizeSum(nums []int, windowSize int) int {
    if windowSize > len(nums) {
        return -1
    }

    sum := 0
    for i := 0; i < windowSize; i += 1 {
        sum += nums[i]
    }

    maxSum := sum
    for i := windowSize; i < len(nums); i += 1 {
        sum = sum + nums[i] - nums[i-windowSize]
        if sum > maxSum {
            maxSum = sum
        }
    }

    return maxSum
}
