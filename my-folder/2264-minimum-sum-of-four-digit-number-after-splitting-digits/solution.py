class Solution:
    def minimumSum(self, num: int) -> int:
        s = sorted('{}'.format(num))
        return 10*(int(s[0]) + int(s[1])) + int(s[2]) + int(s[3])
