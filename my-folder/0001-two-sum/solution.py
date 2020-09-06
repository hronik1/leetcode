class Solution(object):
    def twoSum(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: List[int]
        """
        index = dict()
        for i in range(len(nums)):
            val = nums[i]
            j = index.get(target-val)
            if j is not None:
                return [i, j]
            index[val] = i
        
        return []
        
