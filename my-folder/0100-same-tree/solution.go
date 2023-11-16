/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {
    if p == nil {
        return q == nil
    } else if q == nil {
        return p == nil
    }

    if p.Val != q.Val {
        return false
    }

    if !isSameTree(p.Left, q.Left) {
        return false
    }

    return isSameTree(p.Right, q.Right)
}
