/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func constructMaximumBinaryTree(nums []int) *TreeNode {
    return constructMBT(nums, 0, len(nums)-1)
}

func constructMBT(nums []int, startI int, endI int) *TreeNode {
    if startI > endI {
        return nil
    }
    
    maxI := startI
    for i := startI; i <= endI; i++ {
        if nums[i] > nums[maxI] {
            maxI = i
        }
    }
    
    left := constructMBT(nums, startI, maxI-1)
    right := constructMBT(nums, maxI+1, endI)
    
    return &TreeNode{
        Val: nums[maxI],
        Left: left,
        Right: right,
    }
}
