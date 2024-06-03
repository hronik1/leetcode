class Solution:
    def countKeyChanges(self, s: str) -> int:
        lowered = s.lower()
        return sum([1 if lowered[i] != lowered[i-1] else 0 for i in range(1, len(lowered))])

