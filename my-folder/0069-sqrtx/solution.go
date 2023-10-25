import "fmt"
func mySqrt(x int) int {
    lo := uint64(0)
    hi := uint64(x)
    for lo <= hi {
        mid := lo + (hi-lo)/2
        midSquared := mid * mid

        if midSquared == uint64(x) {
            return int(mid)
        } else if midSquared < uint64(x) {
            if (mid+1) * (mid+1) > uint64(x) {
                return int(mid)
            }
            
            lo = mid + 1
        } else {
            if (mid-1) * (mid-1) < uint64(x) {
                return int(mid-1)
            }
            
            hi = mid - 1
        }
    }
    
    return -1
}
