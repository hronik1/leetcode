/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
    if root == nil {
        return true
    }
    
    curLevel := []*TreeNode{root}
    nonNilChildren := true
    for nonNilChildren {
        nonNilChildren = false
        nextLevel := []*TreeNode{}
        for i, node := range curLevel {
            if node != nil {
                nonNilChildren = true
                symNode := curLevel[len(curLevel)-1-i]
                if symNode == nil || symNode.Val != node.Val {
                    return false
                }
                
                nextLevel = append(nextLevel, []*TreeNode{node.Left, node.Right}...)
            } else {
                nextLevel = append(nextLevel, []*TreeNode{nil, nil}...)
            }
        }
        
        curLevel = nextLevel
    }

    return true
    
    
}
