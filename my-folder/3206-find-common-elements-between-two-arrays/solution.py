class Solution:
    def findIntersectionValues(self, nums1: List[int], nums2: List[int]) -> List[int]:
        s1, s2 = set(nums1), set(nums2)
        return [sum([1 if num1 in s2 else 0 for num1 in nums1]), sum([1 if num2 in s1 else 0 for num2 in nums2])]
