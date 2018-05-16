package main

import (
	"fmt"
	"strings"
)

type MapSum struct {
	HashMap map[string]int
}

/** Initialize your data structure here. */
func Constructor() MapSum {
	return MapSum{make(map[string]int)}
}

func (this *MapSum) Insert(key string, val int) {
	this.HashMap[key] = val
}

func (this *MapSum) Sum(prefix string) int {
	ret := 0
	for k, v := range this.HashMap {
		if strings.Index(k, prefix) == 0 {
			ret += v
		}
	}

	return ret
}

/**
 * Your MapSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(key,val);
 * param_2 := obj.Sum(prefix);
 */
func main() {
	obj := Constructor()
	fmt.Println(obj.Sum("aab"))
	obj.Insert("aab", 33)
	fmt.Println(obj.Sum("aab"))
	fmt.Println(obj.Sum("ab"))
}
