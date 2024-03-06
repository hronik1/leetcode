import string

class Solution:
    def decodeMessage(self, key: str, message: str) -> str:
        mapping = {}
        idx = 0
        for c in key:
            if c != ' ' and c not in mapping:
                mapping[c] = string.ascii_lowercase[idx]
                idx += 1
            if idx >= 26:
                break
        
        return "".join([mapping[c] if c in mapping else c for c in message])

        
