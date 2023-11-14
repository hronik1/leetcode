func removeDuplicates(nums []int) int {
    prev := nums[0]
    count := 1
    writeIndex := 1

    for i := 1; i < len(nums); i += 1 {
        if nums[i] != prev {
            nums[writeIndex] = nums[i]
            writeIndex += 1
            prev = nums[i]
            count = 1
        } else if count < 2 {
            nums[writeIndex] = nums[i]
            writeIndex += 1
            count += 1
        }
    }

    return writeIndex
}
