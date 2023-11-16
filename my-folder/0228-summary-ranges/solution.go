import "fmt"

func summaryRanges(nums []int) []string {
    if len(nums) == 0 {
        return []string{}
    }

    out := []string{}
    curLo := 0
    curHi := 0
    for i := 1; i < len(nums); i += 1 {
        if nums[i] == nums[curHi]+1 {
            curHi = i
        } else {
            out = append(out, formatString(nums, curLo, curHi))
            curLo = i
            curHi = i
        }
    }

    out = append(out, formatString(nums, curLo, curHi))
    
    return out
}

func formatString(nums []int, lo int, hi int) string {
    if lo == hi {
        return fmt.Sprintf("%d", nums[lo])
    } else {
        return fmt.Sprintf("%d->%d", nums[lo], nums[hi])
    }
}
