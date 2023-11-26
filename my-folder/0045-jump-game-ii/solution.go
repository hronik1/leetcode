func jump(nums []int) int {
    count := 0
    i := 0
    for i < len(nums) - 1 {
        if i + nums[i] >= len(nums) - 1 {
            count += 1
            break
        }

        nextI := i
        bestNextJump := i
        for offset := 0; offset <= nums[i] && i + offset < len(nums); offset += 1 {
            nextJump := i + offset + nums[i + offset]
            if nextJump > bestNextJump {
                bestNextJump = nextJump
                nextI = i + offset
            }
        }

        i = nextI
        count += 1
    }
    return count
}
