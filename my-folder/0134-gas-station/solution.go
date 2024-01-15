func canCompleteCircuit(gas []int, cost []int) int {
    diffs := []int{}
    allZero := true
    runningSum := 0
    for i := 0; i < len(gas); i++ {
        diff := gas[i]-cost[i]
        runningSum += diff
        diffs = append(diffs, diff)
        if diff != 0 {
            allZero = false
        }
    }

    if runningSum < 0 {
        return -1
    }

    if allZero {
        return 0
    }
  
    for i := 0; i < len(gas); i++ {
        if diffs[i] > 0 && canCompleteFromI(i, gas, cost) {
            return i
        }
    }

    return -1
}

func canCompleteFromI(start int, gas []int, cost []int) bool {
    runningSum := 0
    for i := start; i < len(gas); i++ {
        runningSum += gas[i]-cost[i]
        if runningSum < 0 {
            return false
        }
    }

    for i := 0; i < start; i++ {
        runningSum += gas[i]-cost[i]
        if runningSum < 0 {
            return false
        }
    }

    return runningSum >= 0
}
