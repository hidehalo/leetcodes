package main

import (
	"../ds"
)

func getDecimalValue(head *ds.ListNode) int {
	ret := 0
	p := head.Next
	for p != nil {
		ret = ret<<1 + p.Val
		p = p.Next
	}

	return ret
}
