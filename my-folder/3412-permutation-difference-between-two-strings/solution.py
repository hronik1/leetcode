class Solution:
    def findPermutationDifference(self, s: str, t: str) -> int:
        s_index = {s[i]: i for i in range(len(s))}
        t_index = {t[i]: i for i in range(len(t))}
        return sum([abs(t_index[c] - s_index[c]) for c in s])
