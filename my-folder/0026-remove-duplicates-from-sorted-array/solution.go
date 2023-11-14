func removeDuplicates(nums []int) int {
    prev := nums[0]
    numUnique := 1
    for i := 1; i < len(nums); i += 1 {
        if nums[i] != prev {
            nums[numUnique] = nums[i]
            numUnique += 1
        }

        prev = nums[i]
    }

    return numUnique
}
