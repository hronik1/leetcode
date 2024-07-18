from collections import defaultdict
class Solution:
    def countKDifference(self, nums: List[int], k: int) -> int:
        counts = defaultdict(int)
        for num in nums:
            counts[num] += 1
        
        out = 0
        for num, count in counts.items():
            out += count * counts.get(num-k, 0)

        return out
        
