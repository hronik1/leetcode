func twoSum(numbers []int, target int) []int {
    i := 0
    j := len(numbers) - 1
    for i < j {
        sum := numbers[i] + numbers[j]
        if sum == target {
            break
        } else if sum < target {
            i += 1
        } else {
            j -= 1
        }
    }

    return []int{i+1, j+1}
}
