class Solution:
    def maximumOddBinaryNumber(self, s: str) -> str:
        l = len(s)
        num_ones = s.count('1')
        return '1'*(num_ones-1) + '0'*(l-num_ones) + '1'
        
