import "slices"

func successfulPairs(spells []int, potions []int, success int64) []int {
    out := []int{}
    if len(spells) == 0 || len(potions) == 0 {
        return out
    }

    slices.Sort(potions)
    for _, v := range spells {
        strength := float64(success)/float64(v)
        index := bstGreaterOrEqual(potions, strength)
        out = append(out, len(potions)-index)
    }

    return out
}

func bstGreaterOrEqual(arr []int, target float64) int {
    if float64(arr[len(arr) - 1]) < target {
        return len(arr)
    }

    lo, hi := 0, len(arr)
    best := len(arr) - 1
    for lo <= hi {
        mid := lo + (hi-lo)/2
        if float64(arr[mid]) >= target {
            best = mid
            hi = mid - 1
        } else {
            lo = mid + 1
        }
    } 

    return best
}
