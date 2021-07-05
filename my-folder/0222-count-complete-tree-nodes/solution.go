/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func countNodes(root *TreeNode) int {
    h := getHeight(root)
    indexLastLeaf := getIndexLastLeaf(root, h)
    
    return int(math.Pow(2.0, float64(h-1))) + indexLastLeaf 
}

func getIndexLastLeaf(root *TreeNode, h int) int {
    lo, hi := 0, int(math.Pow(2.0, float64(h-1)))-1
    originalMid := (hi+1)/2
    lastValidChild := 0
    for lo <= hi {
        mid := lo + (hi-lo)/2
        if isValid(root, originalMid, mid) {
            if mid > lastValidChild {
                lastValidChild = mid 
            }
            lo = mid + 1
        } else {
            hi = mid - 1
        }
    }
    
    return lastValidChild
}

func isValid(root *TreeNode, originalMid int, i int) bool{
    n := root
    for d := originalMid; d > 0; d /= 2 {
        if i/d >= 1 {
            n = n.Right
        } else {
            n = n.Left
        }
        
        i = i%d
    }
    
    return n != nil
}


func getHeight(root *TreeNode) int {
    h := 0
    n := root
    for n != nil {
        n = n.Left
        h++
    }
    
    return h
}
