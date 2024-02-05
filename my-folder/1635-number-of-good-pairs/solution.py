from collections import defaultdict

class Solution:
    def numIdenticalPairs(self, nums: List[int]) -> int:
        d = defaultdict(lambda: 0)
        pairs = 0
        for num in nums:
            pairs += d[num]
            d[num] += 1

        return pairs
