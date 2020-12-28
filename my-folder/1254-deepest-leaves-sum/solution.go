/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func deepestLeavesSum(root *TreeNode) int {
    if root == nil {
        return 0
    }
    
    currLevel := []*TreeNode{root}
    sum := 0
    for len(currLevel) > 0 {
        sum = 0
        nextLevel := []*TreeNode{}
        for _, node := range currLevel {
            sum += node.Val
            
            if node.Left != nil {
                nextLevel = append(nextLevel, node.Left)
            }
            
            if node.Right != nil {
                nextLevel = append(nextLevel, node.Right)
            }
        }
        
        currLevel = nextLevel
    }
    
    return sum
}
