class Solution:
    def semiOrderedPermutation(self, nums: List[int]) -> int:
        one_i = nums.index(1)
        n_i = nums.index(len(nums))
        num_swaps = one_i + (len(nums)-1-n_i)
        if n_i < one_i:
            num_swaps -= 1
        
        return num_swaps
