class Solution:
    def pivotArray(self, nums: List[int], pivot: int) -> List[int]:
        less_than_pivot_count, pivot_count = 0, 0
        for num in nums:
            if num < pivot:
                less_than_pivot_count += 1
            if num == pivot:
                pivot_count += 1
        
        out = [0] * len(nums)
        less_i, pivot_i, greater_i = 0, less_than_pivot_count, less_than_pivot_count+pivot_count
        for num in nums:
            if num < pivot:
                out[less_i] = num
                less_i += 1
            elif num == pivot:
                out[pivot_i] = num
                pivot_i += 1
            else:
                out[greater_i] = num
                greater_i += 1

        return out 
