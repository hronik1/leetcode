class Solution:
    def findAndReplacePattern(self, words: List[str], pattern: str) -> List[str]:
        return [word for word in words if self.matches_pattern(word, pattern)]
    
    def matches_pattern(self, word, pattern):
        if len(word) != len(pattern):
            return False
        
        word_to_pattern = {}
        pattern_to_word = {}
        for i in range(len(word)):
            word_c = word[i]
            pattern_c = pattern[i]
            if word_c not in word_to_pattern:
                if pattern_c in pattern_to_word:
                    return False
                word_to_pattern[word_c] = pattern_c
                pattern_to_word[pattern_c] = word_c

            if word_to_pattern[word_c] != pattern_c:
                return False

        return True
