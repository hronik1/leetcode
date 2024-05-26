class Solution:
    def removeOuterParentheses(self, s: str) -> str:
        count = 0
        out = ''
        for c in s:
            if c == '(':
                if count > 0:
                    out += '('
                count += 1
            else:
                count -= 1
                if count > 0:
                    out += ')'
        return out
        
