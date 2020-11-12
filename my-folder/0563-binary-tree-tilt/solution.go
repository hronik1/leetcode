import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findTilt(root *TreeNode) int {
    out := findTiltHelper(root)
    return out.TiltSum
}

type findTiltHelperOutput struct {
    Sum int
    TiltSum int
}

func findTiltHelper(node *TreeNode) findTiltHelperOutput {
    if node == nil {
        return findTiltHelperOutput{}
    }
    
    leftOut := findTiltHelper(node.Left)
    rightOut := findTiltHelper(node.Right)
    
    tilt := int(math.Abs(float64(leftOut.Sum - rightOut.Sum)))
    
    return findTiltHelperOutput{
        Sum: leftOut.Sum + rightOut.Sum + node.Val,
        TiltSum: tilt + leftOut.TiltSum + rightOut.TiltSum,
    }
}
