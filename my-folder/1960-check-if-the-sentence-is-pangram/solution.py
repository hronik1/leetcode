from collections import Counter
class Solution:
    def checkIfPangram(self, sentence: str) -> bool:
        c = Counter(sentence)
        return len(c) == 26

        
