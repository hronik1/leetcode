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
    
    depth := 1
    maxChildrenDepth := 0
    
    if root.Left != nil {
        maxChildrenDepth = maxDepth(root.Left)
    }
    
    if root.Right != nil {
        rightDepth := maxDepth(root.Right)
        if rightDepth > maxChildrenDepth {
            maxChildrenDepth = rightDepth
        }
    }
    
    return depth + maxChildrenDepth
}
