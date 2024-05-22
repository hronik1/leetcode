class Solution:
    def removeOccurrences(self, s: str, part: str) -> str:
        while True:
            i = s.find(part)
            if i == -1:
                break
            s = s[:i] + s[i+len(part):]
            
        return s
        
