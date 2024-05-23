class Solution:
    def finalString(self, s: str) -> str:
        out = ""
        for c in s:
            if c == 'i':
                out = out[::-1]
            else:
                out += c
        
        return out
