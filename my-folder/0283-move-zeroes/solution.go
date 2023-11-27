func moveZeroes(nums []int)  {
    writeIndex := 0
    for _, v := range nums {
        if v != 0 {
            nums[writeIndex] = v
            writeIndex += 1
        }
    }

    for ; writeIndex < len(nums); writeIndex += 1 {
        nums[writeIndex] = 0
    }
}
