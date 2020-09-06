class Solution(object):
    def twoSum(self, numbers, target):
        """
        :type numbers: List[int]
        :type target: int
        :rtype: List[int]
        """
        i = 1
        j = len(numbers)
        while i < j:
            total = numbers[i-1] + numbers [j-1]
            if total == target:
                return [i,j]
            elif total < target:
                i+=1
            else:
                j-=1
        return []
        
