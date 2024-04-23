class Solution:
    def minPairSum(self, nums: List[int]) -> int:
        sorted_nums = sorted(nums)
        hi = len(sorted_nums) - 1
        return max([sorted_nums[i] + sorted_nums[hi - i] for i in range(int(len(sorted_nums)/2))])
            

        
