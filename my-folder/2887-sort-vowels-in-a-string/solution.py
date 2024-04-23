class Solution:
    def sortVowels(self, s: str) -> str:
        vowels = set(['a', 'A', 'e', 'E', 'i', 'I', 'o', 'O', 'u', 'U'])

        seen_vowels = sorted([c for c in s if c in vowels], key=ord, reverse=True)
        out = ""
        for c in s:
            if c in vowels:
                c = seen_vowels.pop()
            out += c
        
        return out

