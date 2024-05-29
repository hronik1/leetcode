class Solution:
    def differenceOfSums(self, n: int, m: int) -> int:
        return sum([i for i in range(n+1) if i%m != 0]) - sum([i for i in range(n+1) if i%m == 0])
