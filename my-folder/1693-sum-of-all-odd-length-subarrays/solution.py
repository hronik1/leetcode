class Solution:
    def sumOddLengthSubarrays(self, arr: List[int]) -> int:
        out = 0
        for l in range(1, len(arr)+1, 2):
            for i in range(len(arr)+1-l):
                out += sum(arr[i:i+l])
        
        return out
