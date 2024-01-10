/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {
    if root == nil {
        return nil
    }

    invertedLeft := invertTree(root.Left)
    invertedRight := invertTree(root.Right)
    root.Left = invertedRight
    root.Right = invertedLeft

    return root
}
