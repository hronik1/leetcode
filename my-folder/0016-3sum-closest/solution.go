func threeSumClosest(nums []int, target int) int {
    slices.Sort(nums)
    bestSum := nums[0]+nums[1]+nums[len(nums)-1]
    bestDiff := int(math.Abs(float64(target - bestSum)))
    for i := 0; i < len(nums)-2; i++ {
        for j, k := i+1, len(nums)-1; j < k; {
            sum := nums[i]+nums[j]+nums[k]
            diff := target - sum
            if diff < 0 {
                k--
            } else {
                j++
            }

            absDiff := int(math.Abs(float64(diff)))
            if absDiff < bestDiff {
                bestSum = sum
                bestDiff = absDiff
            }
        }
    }

    return bestSum 
}
