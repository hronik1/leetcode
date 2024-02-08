from collections import defaultdict
class Solution:
    def frequencySort(self, s: str) -> str:
        counts = defaultdict(int)
        for c in s:
            counts[c] += 1
        
        sortedCounts = sorted(counts.items(), key=lambda x:x[1], reverse=True)
        out = ''.join([pair[0]*pair[1] for pair in sortedCounts])

        return out
        
