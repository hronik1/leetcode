class Solution:
    def minBitFlips(self, start: int, goal: int) -> int:
        flipped_int = start ^ goal
        return flipped_int.bit_count()
