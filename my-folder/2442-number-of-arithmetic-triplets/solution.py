class Solution:
    def arithmeticTriplets(self, nums: List[int], diff: int) -> int:
        nums_set = set(nums)
        count = 0
        for num in nums:
            j_num = diff + num
            k_num = diff + j_num
            if j_num in nums_set and k_num in nums_set:
                count += 1
        
        return count
