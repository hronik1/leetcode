import math

class Solution:
    def isPowerOfFour(self, n: int) -> bool:
        if n < 1:
            return False

        l = math.log(n, 4)
        return math.floor(l) == math.ceil(l)
        
