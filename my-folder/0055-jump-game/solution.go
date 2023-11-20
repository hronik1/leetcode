func canJump(nums []int) bool {
    leftMostGood := len(nums) - 1
    for i := len(nums) - 2; i >= 0; i -= 1 {
        if i + nums[i] >= leftMostGood {
            leftMostGood = i
        }
    }

    return leftMostGood == 0
}

