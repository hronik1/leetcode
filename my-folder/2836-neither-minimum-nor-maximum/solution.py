class Solution:
    def findNonMinOrMax(self, nums: List[int]) -> int:
        min_num, max_num = min(nums), max(nums)
        for num in nums:
            if num != min_num and num != max_num:
                return num

        return -1
