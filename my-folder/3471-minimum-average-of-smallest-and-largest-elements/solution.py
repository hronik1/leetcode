from collections import deque

class Solution:
    def minimumAverage(self, nums: List[int]) -> float:
        if len(nums) < 2:
            return 0.0

        d = deque(sorted(nums))
        min_item, max_item = d.popleft(), d.pop()
        smallest_avg = (min_item+max_item)/2.0
        while len(d) > 1:
            min_item, max_item = d.popleft(), d.pop()
            avg = (min_item+max_item)/2.0
            if avg < smallest_avg:
                smallest_avg = avg
        
        return smallest_avg



        
