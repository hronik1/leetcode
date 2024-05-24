from collections import deque

class Solution:
    def numberGame(self, nums: List[int]) -> List[int]:
        d = deque(sorted(nums, reverse=True))
        out = []
        while len(d) > 0:
            alice_num = d.pop()
            bob_num = d.pop()
            out.append(bob_num)
            out.append(alice_num)
        
        return out
            
        
