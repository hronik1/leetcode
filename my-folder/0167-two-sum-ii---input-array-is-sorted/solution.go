func twoSum(numbers []int, target int) []int {
    i := 1
    j := len(numbers)
    for i < j {
        sum := numbers[i-1] + numbers[j-1]
        if sum == target {
            return []int{i, j}
        } else if sum < target {
            i+=1
        } else {
            j-=1
        }
    }
    
    return []int{}
}
