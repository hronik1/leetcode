class Solution:
    def arrayPairSum(self, nums: List[int]) -> int:
        sorted_nums = sorted(nums)
        return sum([sorted_nums[2*i] for i in range(math.floor(len(nums)/2))])
