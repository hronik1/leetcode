/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumEvenGrandparent(root *TreeNode) int {
    if root == nil {
        return 0
    }
    
    evenParents := []*TreeNode{}
    oddParents := []*TreeNode{}
    if root.Val%2 == 0 {
        evenParents = append(evenParents, root.Left)
        evenParents = append(evenParents, root.Right)
    } else {
        oddParents = append(oddParents, root.Left)
        oddParents = append(oddParents, root.Right)
    }
    
    sum := 0
    for len(evenParents) > 0 || len(oddParents) > 0 {
        for len(evenParents) > 0 {
            node := evenParents[len(evenParents)-1]
            evenParents = evenParents[:len(evenParents)-1]
            if node == nil {
                continue
            }
            
            if node.Left != nil {
                sum += node.Left.Val
            }
            if node.Right != nil {
                sum += node.Right.Val
            }
            
            if node.Val%2 == 0 {
                evenParents = append(evenParents, node.Left)
                evenParents = append(evenParents, node.Right)
            } else {
                oddParents = append(oddParents, node.Left)
                oddParents = append(oddParents, node.Right)
            } 
        }
        
        for len(oddParents) > 0 {
            node := oddParents[len(oddParents)-1]
            oddParents = oddParents[:len(oddParents)-1]
            if node == nil {
                continue
            }
            
            if node.Val%2 == 0 {
                evenParents = append(evenParents, node.Left)
                evenParents = append(evenParents, node.Right)
            } else {
                oddParents = append(oddParents, node.Left)
                oddParents = append(oddParents, node.Right)
            }  
        }
    }
    
    return sum
}
