/** 
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
    lo := 0
    hi := n
    
    for lo <= hi {
        mid := lo + (hi-lo)/2
        g := guess(mid)
        if g == 0 {
            return mid
        } else if g < 1 {
            hi = mid - 1
        } else {
            lo = mid + 1
        }
    }
    
    return -1
}
