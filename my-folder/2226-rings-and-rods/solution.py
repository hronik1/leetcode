from collections import defaultdict
class Solution:
    def countPoints(self, rings: str) -> int:
        seen = defaultdict(set)
        for i in range(int(len(rings)/2)):
            index = 2*i
            color = rings[index]
            ring = rings[index+1]
            seen[ring].add(color)
        
        count = 0
        for ring, colors in seen.items():
            if len(colors) == 3:
                count += 1
        
        return count
        
