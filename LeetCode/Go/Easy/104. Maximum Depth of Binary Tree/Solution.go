/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
    if root == nil {
		return 0
	}
	// fmt.Println(root.Val) 
	levelLeft := maxDepth(root.Left)
	levelRight := maxDepth(root.Right)

	return max(levelLeft, levelRight) + 1
}