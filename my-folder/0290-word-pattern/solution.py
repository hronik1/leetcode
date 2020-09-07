class Solution(object):
    def wordPattern(self, pattern, strs):
        """
        :type pattern: str
        :type str: str
        :rtype: bool
        """
        words = strs.split(" ")
        if len(pattern) != len(words):
            return False
        
        mapping = dict()
        reverse_mapping = dict()
        for i in range(len(pattern)):
            w = mapping.get(pattern[i])
            if w and w != words[i]:
                return False
            
            c = reverse_mapping.get(words[i])
            if c and c != pattern[i]:
                return False
            
            mapping[pattern[i]] = words[i]
            reverse_mapping[words[i]] = pattern[i]
        
        return True
