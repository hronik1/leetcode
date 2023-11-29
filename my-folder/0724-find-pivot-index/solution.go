func pivotIndex(nums []int) int {
    sumBefore := []int {0}
    for i := 1; i < len(nums); i += 1 {
        sumBefore = append(sumBefore, sumBefore[i-1] + nums[i-1])
    }

    sumAfter := make([]int, len(nums), len(nums))
    sumAfter[len(nums) - 1] = 0
    for i := len(nums) - 2; i >= 0; i -= 1 {
        sumAfter[i] = sumAfter[i+1] + nums[i+1]
    }

    pivotIndex := -1
    for i := 0; i < len(nums); i += 1 {
        if sumBefore[i] == sumAfter[i] {
            pivotIndex = i
            break
        }
    }

    return pivotIndex
}
