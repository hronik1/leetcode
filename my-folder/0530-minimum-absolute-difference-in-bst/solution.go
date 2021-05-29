/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func getMinimumDifference(root *TreeNode) int {
    out := getMinimumDifferenceHelper(root)
    return out.MinDiff
}

type DiffHelperData struct {
    MinVal int
    MaxVal int
    MinDiff int
}

func getMinimumDifferenceHelper(root *TreeNode) DiffHelperData {
    out := DiffHelperData{
        MinVal: root.Val,
        MaxVal: root.Val,
        MinDiff: -1,
    }
    
    if root.Left != nil {
        left := getMinimumDifferenceHelper(root.Left)
        out.MinVal = left.MinVal
        out.MinDiff = root.Val - left.MaxVal
        if left.MinDiff > -1 && left.MinDiff < out.MinDiff {
            out.MinDiff = left.MinDiff
        }
    }
    
    if root.Right != nil {
        right := getMinimumDifferenceHelper(root.Right)
        out.MaxVal = right.MaxVal
        if out.MinDiff == -1 || (right.MinVal - root.Val) < out.MinDiff {
            out.MinDiff = right.MinVal - root.Val
        }
        if right.MinDiff > -1 && right.MinDiff < out.MinDiff {
            out.MinDiff = right.MinDiff
        }
    }
    
    return out
}
