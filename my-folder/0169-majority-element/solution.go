//import "slices"
func majorityElement(nums []int) int {
    leader := nums[0]
    count := 1
    for i := 1; i < len(nums); i += 1 {
        if nums[i] == leader {
            count += 1
        } else {
            count -= 1
        }

        if count == 0 {
            leader = nums[i]
            count = 1 // is this right?
        }
    }

    return leader
}
