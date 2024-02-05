class Solution:
    def firstUniqChar(self, s: str) -> int:
        seen = set()
        nonUnique = set()
        for c in s:
            if c in seen:
                nonUnique.add(c)
            seen.add(c)

        for i in range (len(s)):
            if s[i] not in nonUnique:
                return i

        return -1

        
