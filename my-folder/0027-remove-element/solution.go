func removeElement(nums []int, val int) int {
    readI := 0
    writeI := 0
    for readI < len(nums) {
        if nums[readI] != val {
            nums[writeI] = nums[readI]
            writeI += 1
        }

        readI += 1
    }
    return writeI
}
