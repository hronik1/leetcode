func increasingTriplet(nums []int) bool {
    smallestBefore := make([]int, len(nums), len(nums))
    largestAfter := make([]int, len(nums), len(nums))

    smallestBefore[0] = nums[0]
    for i := 1; i < len(nums); i++ {
        smallest := nums[i]
        if smallestBefore[i-1] < smallest {
            smallest = smallestBefore[i-1]
        }

        smallestBefore[i] = smallest
    }

    largestAfter[len(nums)-1] = nums[len(nums)-1]
    for i := len(nums)-2; i >= 0; i-- {
        largest := nums[i]
        if largestAfter[i+1] > largest {
            largest = largestAfter[i+1]
        }

        largestAfter[i] = largest
    }

    for i := 0; i < len(nums); i++ {
        if smallestBefore[i] < nums[i] && nums[i] < largestAfter[i] {
            return true
        }
    }

    return false
}
