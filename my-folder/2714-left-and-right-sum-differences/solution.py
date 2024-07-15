class Solution:
    def leftRightDifference(self, nums: List[int]) -> List[int]:
        out = []
        left_sum, right_sum = 0, sum(nums[1:])
        for i in range(len(nums)):
            out.append(abs(right_sum-left_sum))
            left_sum += nums[i]
            if i < len(nums)-1:
                right_sum -= nums[i+1]

        return out
