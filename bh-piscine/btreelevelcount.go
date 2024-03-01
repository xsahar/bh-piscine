package piscine

func BTreeLevelCount(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftHeight := BTreeLevelCount(root.Left)
	rightHeight := BTreeLevelCount(root.Right)

	// Return the maximum height between the left and right subtrees, plus 1 for the current level
	if leftHeight > rightHeight {
		return leftHeight + 1
	}
	return rightHeight + 1
}
