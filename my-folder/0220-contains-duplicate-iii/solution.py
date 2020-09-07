from sortedcontainers import SortedSet

class Solution(object):
    def containsNearbyAlmostDuplicate(self, nums, k, t):
        """
        :type nums: List[int]
        :type k: int
        :type t: int
        :rtype: bool
        """
        if k <= 0:
            return False
        
        ss = SortedSet([(nums[i], i) for i in range(min(len(nums), k))])
        for i in range(len(nums)):
            ss.discard((nums[i], i))
            if i+k < len(nums):
                ss.add((nums[i+k], i+k))
                
            lo = 0
            hi = len(ss)-1 
            while lo <= hi:
                mid = lo + (hi-lo)/2
                if abs(ss[mid][0] - nums[i]) <= t:
                    return True
                elif ss[mid][0] > nums[i]: 
                    hi = mid - 1 
                else:
                    lo = mid + 1
        
        return False
