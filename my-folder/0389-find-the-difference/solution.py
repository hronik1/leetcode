from collections import defaultdict
class Solution:

    def findTheDifference(self, s: str, t: str) -> str:
        d = defaultdict(int)
        for c in t:
            d[c] += 1
        
        for c in s:
            d[c] -= 1
            if d[c] == 0:
                del d[c]
        
        # assert length of d is 1
        return list(d.keys())[0]
        
