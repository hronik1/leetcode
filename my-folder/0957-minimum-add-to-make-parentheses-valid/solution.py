from collections import deque
class Solution:
    def minAddToMakeValid(self, s: str) -> int:
        d = deque()
        for c in s:
            if c == '(':
                d.append(c)
            if c == ')':
                if len(d) > 0 and d[-1] == '(':
                    d.pop()
                else:
                    d.append(c)

        return len(d)        
