package piscine

func BTreeMax(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	// Traverse to the rightmost node, which will have the maximum value
	for root.Right != nil {
		root = root.Right
	}

	return root
}
