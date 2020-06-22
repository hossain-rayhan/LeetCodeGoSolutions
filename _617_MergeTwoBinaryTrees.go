/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
    if(t1 == nil && t2 == nil){
        return nil;
    }    
    
    if(t1 != nil && t2 != nil){
        n := TreeNode{};
        n.Val = t1.Val + t2.Val;
        n.Left = mergeTrees(t1.Left, t2.Left);
        n.Right = mergeTrees(t1.Right, t2.Right);
        return &n;
    }else if (t1 == nil){
        return t2;
    }else {
        return t1;
    }
}
