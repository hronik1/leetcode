import "slices"

func trap(height []int) int {
    if len(height) <= 2 {
        return 0
    }

    before := maxBefore(height)
    slices.Reverse(height)
    after := maxBefore(height)
    slices.Reverse(after)
    slices.Reverse(height)

    trapped := 0
    for i, v := range height {
        curHeight := before[i]
        if after[i] < curHeight {
            curHeight = after[i]
        }

        curHeight -= v
        if curHeight > 0 {
            trapped += curHeight
        }
    }

    return trapped
}

func maxBefore(height []int) []int {
    ret := []int{0}
    for i := 1; i < len(height); i += 1 {
        curMax := height[i-1]
        if ret[i-1] > curMax {
            curMax = ret[i-1]
        }

        ret = append(ret, curMax)
    }

    return ret
}
