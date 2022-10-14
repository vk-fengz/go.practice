package _5_unique_binary_search_trees_ii

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}
	return buildBST(1, n)
}

func buildBST(lo, hi int) (res []*TreeNode) {
	if lo > hi {
		res = append(res, nil)
		return res
	}

	for i := lo; i <= hi; i++ {
		leftTree := buildBST(lo, i-1)
		rightTree := buildBST(i+1, hi)

		for _, left := range leftTree {
			for _, right := range rightTree {
				root := &TreeNode{Val: i,
					Left: left, Right: right}
				res = append(res, root)
			}
		}
	}

	return res
}
