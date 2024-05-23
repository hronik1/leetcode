class Solution:
    def sortSentence(self, s: str) -> str:
        reversed_sorted = sorted([word[::-1] for word in s.split(" ")])
        return " ".join(word[:0:-1] for word in reversed_sorted)
