class Solution:
    def countDigits(self, num: int) -> int:
        count = 0
        digits = num
        while digits > 0:
            digit = digits % 10
            if num % digit == 0:
                count += 1

            digits = int(digits/10)
        return count
