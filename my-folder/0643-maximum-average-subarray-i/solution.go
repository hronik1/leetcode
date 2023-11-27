func findMaxAverage(nums []int, k int) float64 {
    sum := 0.0
    for i := 0; i < k; i += 1 {
        sum += float64(nums[i])
    }

    maxAvg := sum/float64(k)
    for i := k; i < len(nums); i += 1 {
        sum += float64(nums[i])
        sum -= float64(nums[i-k])
        avg := sum/float64(k)
        if avg > maxAvg {
            maxAvg = avg
        }
    }

    return maxAvg
}
