import "math"
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxAncestorDiff(root *TreeNode) int {
    if root == nil {
        return 0
    }
    
    out := maxAncestorDiffHelper(root)
    return out.MaxDiff
}

type maxAncestorDiffHelperOutput struct {
    MaxDiff int
    MaxVal int
    MinVal int
}

func maxAncestorDiffHelper(root *TreeNode) maxAncestorDiffHelperOutput {
    out := maxAncestorDiffHelperOutput{}
    // shouldn't happen but safeguarding
    if root == nil {
        return out
    }
    
    out.MaxVal = root.Val
    out.MinVal = root.Val
    
    if root.Left != nil {
        leftOut := maxAncestorDiffHelper(root.Left)
        if leftOut.MaxDiff > out.MaxDiff {
            out.MaxDiff = leftOut.MaxDiff
        }
        
        diff := int(math.Abs(float64(root.Val - leftOut.MaxVal)))
        if diff > out.MaxDiff {
            out.MaxDiff = diff
        }
                    
        diff = int(math.Abs(float64(root.Val - leftOut.MinVal)))
        if diff > out.MaxDiff {
            out.MaxDiff = diff
        }
                   
        if leftOut.MaxVal > out.MaxVal {
            out.MaxVal = leftOut.MaxVal              
        }
        
        if leftOut.MinVal < out.MinVal {
            out.MinVal = leftOut.MinVal              
        }
    }
    
    if root.Right != nil {
        rightOut := maxAncestorDiffHelper(root.Right)
        if rightOut.MaxDiff > out.MaxDiff {
            out.MaxDiff = rightOut.MaxDiff
        }
        
        diff := int(math.Abs(float64(root.Val - rightOut.MaxVal)))
        if diff > out.MaxDiff {
            out.MaxDiff = diff
        }
                    
        diff = int(math.Abs(float64(root.Val - rightOut.MinVal)))
        if diff > out.MaxDiff {
            out.MaxDiff = diff
        }
                   
        if rightOut.MaxVal > out.MaxVal {
            out.MaxVal = rightOut.MaxVal              
        }
        
        if rightOut.MinVal < out.MinVal {
            out.MinVal = rightOut.MinVal              
        }
    }
    
    return out
}
