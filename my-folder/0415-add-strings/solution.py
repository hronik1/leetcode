class Solution:
    def addStrings(self, num1: str, num2: str) -> str:
        digits = []
        carry = 0
        reversed_num1 = num1[::-1]
        reversed_num2 = num2[::-1]
        longest_length = max(len(num1), len(num2))
        for i in range(longest_length):
            digit1 = self.get_digit(reversed_num1, i)
            digit2 = self.get_digit(reversed_num2, i)
            s = carry + digit1 + digit2
            digits.append(s%10)
            carry = int(s/10)

        if carry > 0:
            digits.append(carry)

        return "".join(str(digit) for digit in reversed(digits))
    
    def get_digit(self, num, i):
        return int(num[i]) if i < len(num) else 0
        
