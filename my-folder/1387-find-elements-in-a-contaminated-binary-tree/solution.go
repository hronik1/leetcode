/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type FindElements struct {
    Root *TreeNode
}


func Constructor(root *TreeNode) FindElements {
    if root == nil {
        return FindElements {}
    }
    
    cleanse(root, 0)
    return FindElements{
        Root: root,
    }
}


func (this *FindElements) Find(target int) bool {
    return find(this.Root, target)
}

func cleanse(root *TreeNode, val int) {
    root.Val = val
    if root.Left != nil {
        cleanse(root.Left, 2*val+1)
    }
    if root.Right != nil {
        cleanse(root.Right, 2*val+2)
    }
}

func find(root *TreeNode, target int) bool {
    if root == nil {
        return false
    }
    
    if root.Val == target {
        return true
    }
    
    if find(root.Left, target) {
       return true 
    }
    
    return find(root.Right, target)
}
/**
 * Your FindElements object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Find(target);
 */
