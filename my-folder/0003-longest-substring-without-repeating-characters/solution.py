class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        if not len(s):
            return 0

        best_len = 1
        start = 0
        index = {s[0]: 0}
        for end in range(1, len(s)):
            c = s[end]
            if c in index and index[c] >= start:
                start = index[c] + 1
            
            index[c] = end
            cur_len = end - start + 1
            if best_len < cur_len:
                best_len = cur_len
        
        return best_len
