class Solution:
    def find132pattern(self, nums: List[int]) -> bool:
        s = []
        third = None
        for i in reversed(range(len(nums))):
            if (third is not None and nums[i] < third):
                return True
            
            while len(s) > 0 and s[-1] < nums[i]:
                third = s[-1]
                s.pop()
            
            s.append(nums[i])
        
        return False
