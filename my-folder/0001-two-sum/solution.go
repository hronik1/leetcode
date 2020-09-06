func twoSum(nums []int, target int) []int {
    index := map[int]int{}
    for i, val := range(nums) {
        if j, ok := index[target-val]; ok {
            return []int{i, j}
        }
        index[val] = i
    }
    
    return []int{}
}
