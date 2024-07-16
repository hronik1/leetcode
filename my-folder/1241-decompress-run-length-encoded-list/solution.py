class Solution:
    def decompressRLElist(self, nums: List[int]) -> List[int]:
        out = []
        for i in range(int(len(nums)/2)):
            offset = 2*i
            out.extend([nums[offset+1]] * nums[offset])

        return out
