import "slices"

func maxOperations(nums []int, k int) int {
    operations := 0
    slices.Sort(nums)
    for i, j := 0, len(nums) - 1; i < j; {
        sum := nums[i] + nums[j] 
        if sum == k {
            operations += 1
            i += 1
            j -= 1
        } else if sum < k {
            i += 1
        } else {
            j -= 1
        }
    }

    return operations
}
