/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect(root *Node) *Node {
    if root == nil {
        return root
    }
    
    level := []*Node{root}
    for len(level) > 0 {
        newLevel := []*Node{}
        for i, node := range level {
            if i < len(level) - 1 {
                node.Next = level[i+1]
            }
            
            if node.Left != nil {
                newLevel = append(newLevel, node.Left)
            }
            
            if node.Right != nil {
                newLevel = append(newLevel, node.Right)
            }
        }
        
        level = newLevel
    }
    
    return root
}
