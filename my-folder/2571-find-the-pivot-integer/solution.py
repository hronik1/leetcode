class Solution:
    def pivotInteger(self, n: int) -> int:
        arr = list(range(n+1))
        lo, hi = 1, n
        while lo <= hi:
            mid = lo + int((hi-lo)/2)
            diff = sum(arr[mid:n+1]) - sum(arr[1:mid+1])
            if diff > 0:
                lo = mid + 1
            elif diff < 0:
                hi = mid - 1
            else:
                return mid

        return -1

