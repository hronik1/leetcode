class Solution:
    def divideArray(self, nums: List[int], k: int) -> List[List[int]]:
        nums.sort()
        groups = []
        groupNums = len(nums)/3
        for groupNum in range(int(groupNums)):
            group = []
            prev = nums[3*groupNum]
            for i in range(3):
                cur = nums[3*groupNum+i]
                if cur - prev > k:
                    return []
                group.append(cur)
                
            groups.append(group)

        
        return groups
        
