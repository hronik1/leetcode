class Solution:
    def buildArray(self, nums: List[int]) -> List[int]:
        out = []
        for v in nums:
            out.append(nums[v])

        return out
        
