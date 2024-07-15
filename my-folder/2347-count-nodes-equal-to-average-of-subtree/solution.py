from collections import namedtuple
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
Subresult = namedtuple('Subresult', ['size', 'running_sum', 'matches'])

class Solution:
    def averageOfSubtree(self, root: TreeNode) -> int:
        subresult = self.averageOfSubtreeHelper(root)
        return subresult.matches
    
    def averageOfSubtreeHelper(self, root):
        if root is None:
            return Subresult(0, 0, 0)

        left_subresult = self.averageOfSubtreeHelper(root.left)
        right_subresult = self.averageOfSubtreeHelper(root.right)

        size = 1 + left_subresult.size + right_subresult.size
        running_sum = root.val + left_subresult.running_sum + right_subresult.running_sum
        matches = left_subresult.matches + right_subresult.matches
        if int(running_sum/size) == root.val:
            matches += 1
        
        return Subresult(size, running_sum, matches)


