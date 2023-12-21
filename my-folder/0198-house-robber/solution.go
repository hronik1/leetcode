func rob(nums []int) int {
    if len(nums) == 1 {
        return nums[0]
    }

    memo := make([]int, len(nums), len(nums))
    for i := 0; i < len(memo); i++ {
        memo[i] = -1
    }
    best := robStartingAtIth(nums, 0, memo)
    if skipped := robStartingAtIth(nums, 1, memo); skipped > best {
        best = skipped
    }

    return best
}

func robStartingAtIth(nums []int, i int, memo []int) int {
    if memo[i] > -1 {
        return memo[i]
    }

    if i >= len(nums)-2 {
        return nums[i]
    }

    total := nums[i]
    best := robStartingAtIth(nums, i+2, memo)
    if i < len(nums)-3 {
        if skipped := robStartingAtIth(nums, i+3, memo); skipped > best {
            best = skipped
        }
    }

    total += best
    memo[i] = total

    return total

}
