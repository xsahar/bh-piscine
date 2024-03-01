package piscine

func BTreeIsBinary(root *TreeNode) bool {
	return isBinarySearchTree(root, nil, nil)
}

func isBinarySearchTree(node *TreeNode, min, max *TreeNode) bool {
	if node == nil {
		return true
	}

	// Check if the node's data violates the binary search tree property
	if (min != nil && node.Data <= min.Data) || (max != nil && node.Data >= max.Data) {
		return false
	}

	// Recursively check the left and right subtrees
	return isBinarySearchTree(node.Left, min, node) && isBinarySearchTree(node.Right, node, max)
}
