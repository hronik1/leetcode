func dailyTemperatures(temperatures []int) []int {
    out := make([]int, len(temperatures), len(temperatures))
    seenTempIndex := [][]int{}
    for i, temp := range temperatures {
        for len(seenTempIndex) > 0 && seenTempIndex[len(seenTempIndex)-1][0] < temp {
            top := seenTempIndex[len(seenTempIndex)-1]
            seenTempIndex = seenTempIndex[:len(seenTempIndex)-1]
            out[top[1]] = i - top[1]
        }

        seenTempIndex = append(seenTempIndex, []int{temp, i})
    }
    return out
}
