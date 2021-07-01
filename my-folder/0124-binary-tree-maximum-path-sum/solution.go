/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxPathSum(root *TreeNode) int {
    if root == nil {
        return 0
    }
    
    out := helper(root)
    return out.CurBestPath
}

func helper(root *TreeNode) HelperOut {
    out := HelperOut{
        NodeToBestDescendant: root.Val,
    }
    
    leftVisited := false
    var leftOut HelperOut
    if root.Left != nil {
        leftVisited = true
        leftOut = helper(root.Left)  
    }
    
    rightVisited := false
    var rightOut HelperOut
    if root.Right != nil {
        rightVisited = true
        rightOut = helper(root.Right)
    }
    
    if leftOut.NodeToBestDescendant > rightOut.NodeToBestDescendant && leftOut.NodeToBestDescendant > 0 {
        out.NodeToBestDescendant += leftOut.NodeToBestDescendant
    } else if rightOut.NodeToBestDescendant > 0 {
        out.NodeToBestDescendant += rightOut.NodeToBestDescendant
    }
    
    curPath := root.Val
    if leftOut.NodeToBestDescendant > 0 {
        curPath += leftOut.NodeToBestDescendant
    }
    if rightOut.NodeToBestDescendant > 0 {
        curPath += rightOut.NodeToBestDescendant
    }
    
    curBestPath := curPath
    if leftVisited && leftOut.CurBestPath > curBestPath {
        curBestPath = leftOut.CurBestPath
    }
    if rightVisited && rightOut.CurBestPath > curBestPath {
        curBestPath = rightOut.CurBestPath
    }
    
    out.CurBestPath = curBestPath
    return out
}

type HelperOut struct {
    CurBestPath int
    NodeToBestDescendant int
}


