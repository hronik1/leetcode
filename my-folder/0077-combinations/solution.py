from itertools import combinations

class Solution:
    def combine(self, n: int, k: int) -> List[List[int]]:
        nums = list(range(1, n+1))
        return list(combinations(nums, k))
