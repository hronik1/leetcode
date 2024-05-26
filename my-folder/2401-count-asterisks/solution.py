class Solution:
    def countAsterisks(self, s: str) -> int:
        count = 0
        in_bar = False
        for c in s:
            if c == '*' and not in_bar:
                count += 1
            
            if c == '|':
                in_bar = not in_bar

        return count
        
