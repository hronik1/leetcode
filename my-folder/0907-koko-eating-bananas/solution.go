import "fmt"
import "math"
func minEatingSpeed(piles []int, h int) int {
    lo := 1
    hi := piles[0]
    for _, v := range piles {
        if v > hi {
            hi = v
        }
    }

    minSpeed := hi
    for lo <= hi {
        mid := lo + (hi-lo)/2
        if canEat(piles, h, mid) {
            minSpeed = mid
            hi = mid - 1
        } else {
            lo = mid + 1
        }
    }

    return minSpeed
}

func canEat(piles []int, h int, k int) bool {
    d := 0
    for _, v := range piles {
        d += int(math.Ceil(float64(v)/float64(k)))
    }

    return d <= h
}
