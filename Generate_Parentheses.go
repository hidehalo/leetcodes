// @see https://leetcode.com/problems/generate-parentheses/discuss/139455/Clean-JavaScript-backtracking-solution
// function generateParenthesis(n) {
// 	let res = [];
// 	generate(n, n, '', res);
// 	return res;
//   }

//   function generate(l, r, s, res) { // l: left remaining, r: right remaining
// 	if (l > r) return;  // e.g. ))(

// 	if (!l && !r) return res.push(s);

// 	if (l) generate(l - 1, r, s + '(', res);
// 	if (r) generate(l, r - 1, s + ')', res);
//   }
package main

import (
	"fmt"
)

func generateParenthesis(n int) []string {
	ret := make([]string, 0, 1000)
	bs(n, n, "", &ret)

	return ret
}

func bs(l int, r int, s string, ret *[]string) {
	if l > r {
		return
	}
	if l <= 0 && r <= 0 {
		*ret = append(*ret, s)
	}
	if l > 0 {
		bs(l-1, r, s+"(", ret)
	}
	if r > 0 {
		bs(l, r-1, s+")", ret)
	}
}

func main() {
	fmt.Println(generateParenthesis(3))
}
