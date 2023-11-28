func largestAltitude(gain []int) int {
    highest := 0
    cur := 0
    for _, v := range gain {
        cur += v
        if cur > highest {
            highest = cur
        }
    }

    return highest
}
