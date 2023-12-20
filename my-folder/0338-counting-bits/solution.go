import "strconv"

func countBits(n int) []int {
    counts := []int{}
    for i := 0; i <= n; i++ {
        counts = append(counts, count1s(i))
    }

    return counts
}

func count1s(n int) int {
    count := 0
    bin := strconv.FormatInt(int64(n), 2)
    for _, v := range bin {
        if v == '1' {
            count++
        }
    }

    return count
}
