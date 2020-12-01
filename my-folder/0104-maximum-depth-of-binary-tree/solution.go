/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    
    maxChildDepth := maxDepth(root.Left)
    rightDepth := maxDepth(root.Right)
    if rightDepth > maxChildDepth {
        maxChildDepth = rightDepth
    }
    
    return 1 + maxChildDepth
}
