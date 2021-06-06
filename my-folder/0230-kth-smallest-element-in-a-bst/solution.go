/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
    out := helper(root, k, 0)
    return out.Val
}

type helperOut struct {
    SmallerCount int
    Val int
}

func helper(root *TreeNode, k int, smaller int) helperOut {
    out := helperOut{
        SmallerCount: 0,
        Val: -1,
    }
    
    if root.Left != nil {
        res := helper(root.Left, k, smaller)
        if res.Val > -1 {
            out.Val = res.Val
            return out
        }
        
        out.SmallerCount += res.SmallerCount //double check this
    }
    
    out.SmallerCount++
    if out.SmallerCount + smaller == k {
        out.Val = root.Val
        return out
    }
    
    if root.Right != nil {
        res := helper(root.Right, k, out.SmallerCount+smaller)
        if res.Val > -1 {
            out.Val = res.Val
            return out
        }
        
        out.SmallerCount += res.SmallerCount //double check this
    }
    
    return out
}
