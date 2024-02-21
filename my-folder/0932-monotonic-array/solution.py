class Solution:
    def isMonotonic(self, nums: List[int]) -> bool:
        if len(nums) == 0:
            return True

        prev = nums[0]
        increasing = None
        for i in range(1, len(nums)):
            if increasing is None:
                if nums[i] < prev:
                    increasing = True
                elif nums[i] > prev:
                    increasing = False
            else:
                if nums[i] < prev and not increasing:
                    return False
                elif nums[i] > prev and increasing:
                    return False                
            
            prev = nums[i]

        return True
