class Solution:
    def findThePrefixCommonArray(self, A: List[int], B: List[int]) -> List[int]:
        seenA = set()
        seenB = set()
        C = []
        for i in range(len(A)):
            seenA.add(A[i])
            seenB.add(B[i])
            C.append(len(seenA & seenB))
        return C
