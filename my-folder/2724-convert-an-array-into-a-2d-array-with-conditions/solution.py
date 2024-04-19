class Solution:
    def findMatrix(self, nums: List[int]) -> List[List[int]]:
        counts = {}
        for num in nums:
            counts[num] = counts.get(num, 0) + 1

        max_count = max(counts.values())
        out = [[] for i in range(max_count)]
        for num, count in counts.items():
            for i in range(count):
                out[i].append(num)

        return out
