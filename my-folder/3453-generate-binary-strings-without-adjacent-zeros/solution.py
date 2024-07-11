class Solution:
    def validStrings(self, n: int) -> List[str]:
        return self.validStringsHelper(n, False) + self.validStringsHelper(n, True)

    def validStringsHelper(self, n: int, has_leading_one: bool) -> List[str]:
        if n == 1:
            if has_leading_one:
                return ["1"]
            else:
                return ["0"]
        
        substrings = self.validStringsHelper(n-1, True)
        if has_leading_one:
            substrings += self.validStringsHelper(n-1, False)
        
        digit_to_append = "0"
        if has_leading_one:
            digit_to_append = "1"
        
        return [digit_to_append + substring for substring in substrings]
