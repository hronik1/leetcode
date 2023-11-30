/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func goodNodes(root *TreeNode) int {
    if root == nil {
        return 0
    }

    return goodNodesHelper(root, root.Val)
}

func goodNodesHelper(root *TreeNode, maxVal int) int {
    if root == nil {
        return 0
    }

    count := 0
    if root.Val >= maxVal {
        count += 1
        maxVal = root.Val
    }

    count += goodNodesHelper(root.Left, maxVal)
    count += goodNodesHelper(root.Right, maxVal)

    return count
}
