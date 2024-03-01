package piscine

func BTreeMin(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	// Traverse to the leftmost node, which will have the minimum value
	for root.Left != nil {
		root = root.Left
	}

	return root
}
