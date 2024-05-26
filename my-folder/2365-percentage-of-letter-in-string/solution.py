from collections import Counter
class Solution:
    def percentageLetter(self, s: str, letter: str) -> int:
        counter = Counter(s)
        return int(100*(counter.get(letter,0)/len(s)))
