/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
    leaves1 := getLeaves(root1)
    leaves2 := getLeaves(root2)
    if len(leaves1) != len(leaves2) {
        return false
    }

    for i, _ := range leaves1 {
        if leaves1[i] != leaves2[i] {
            return false
        }
    }

    return true
}

func getLeaves(root *TreeNode) []int {
    out := []int{}
    if root.Left == nil && root.Right == nil {
        out = append(out, root.Val) 
    }

    if root.Left != nil {
        out = append(out, getLeaves(root.Left)...)
    }

    if root.Right != nil {
        out = append(out, getLeaves(root.Right)...)
    }

    return out
}
