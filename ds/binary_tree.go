package ds

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func insert(r *TreeNode, n int) {
	if r == nil {
		return
	}
	p := r
	var q *TreeNode
	for p != nil {
		q = p
		if p.Val < n {
			p = p.Right
		} else {
			p = p.Left
		}
	}
	p = &TreeNode{
		n,
		nil,
		nil,
	}
	if q.Val <= n {
		q.Right = p
	} else {
		q.Left = p
	}
}

func newBST(nums []int) *TreeNode {
	r := &TreeNode{nums[0], nil, nil}
	for i := 1; i < len(nums); i++ {
		insert(r, nums[i])
	}
	return r
}
