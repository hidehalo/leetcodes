package ds

import (
	"fmt"
)

type Node struct {
	Check string
	// Hash table version
	Branches map[string]*Node
}

func Constructor(ss []string) *Node {
	root := &Node{
		Check:    "",
		Branches: make(map[string]*Node),
	}
	for _, s := range ss {
		root.Insert(s)
	}

	return root
}

func (root *Node) Insert(s string) {
	s += "$"
	size := len(s)
	p := root
	for i := 0; i < size; i++ {
		key := string(s[i])
		if _, ok := p.Branches[key]; !ok {
			p.Branches[key] = &Node{
				key,
				make(map[string]*Node),
			}
			p = p.Branches[key]
			continue
		}
		p = p.Branches[key]
	}
}

func (root *Node) Query(s string) bool {
	size := len(s)
	p := root
	for i := 0; i < size; i++ {
		key := string(s[i])
		if _, ok := p.Branches[key]; !ok {
			return false
		}
		p = p.Branches[key]
	}
	if p.Branches["$"] != nil {
		return true
	}

	return false
}

func (root *Node) Travasle() {
	if root == nil {
		return
	}
	for _, b := range root.Branches {
		fmt.Println(b.Check)
		b.Travasle()
	}
}

// func main() {
// 	root := Constructor([]string{"anne", "ane", "an"})
// 	root.Travasle()
// 	samples := []string{
// 		"an", "ann", "anne", "ane", "anan", "anen", "n", "",
// 	}
// 	for _, s := range samples {
// 		var ret string
// 		if root.Query(s) {
// 			ret = "Yes"
// 		} else {
// 			ret = "No"
// 		}
// 		fmt.Printf("Query String \"%s\" Result: %s\n", s, ret)
// 	}
// }
