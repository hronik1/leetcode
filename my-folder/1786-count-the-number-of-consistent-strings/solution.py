class Solution:
    def countConsistentStrings(self, allowed: str, words: List[str]) -> int:
        allowed_set = set(allowed)
        return len([word for word in words if self.are_all_allowed(allowed_set, word)])

    def are_all_allowed(self, allowed_set, word):
        for c in word:
            if c not in allowed_set:
                return False
        return True
