package main

type Slot struct {
	Key  int
	Val  int
	Next *Slot
}

// MyHashMap is a bad hash map implement
// no table doubling lol
type MyHashMap struct {
	store []*Slot
}

func (this *MyHashMap) hash(key int) int {
	return key % len(this.store)
}

/** Initialize your data structure here. */
func Constructor() MyHashMap {
	return MyHashMap{
		store: make([]*Slot, 1024),
	}
}

/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int) {
	hash := this.hash(key)
	newSlot := &Slot{
		Key:  key,
		Val:  value,
		Next: nil,
	}
	if this.store[hash] == nil {
		this.store[hash] = newSlot
	} else {
		slots := this.store[hash]
		p := slots
		for p != nil {
			if p.Key == key {
				p.Val = value
				break
			} else if p.Next == nil {
				p.Next = newSlot
				break
			}
			p = p.Next
		}
	}
}

/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
	slots := this.store[this.hash(key)]
	if slots == nil {
		return -1
	}
	p := slots
	for p != nil {
		if p.Key == key {
			return p.Val
		}
		p = p.Next
	}

	return -1
}

/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int) {
	slots := this.store[this.hash(key)]
	if slots == nil {
		return
	}
	q, p := slots, slots
	for p != nil {
		if p.Key == key {
			if q == p {
				this.store[this.hash(key)] = p.Next
			} else {
				q.Next = p.Next
			}
			break
		}
		q = p
		p = p.Next
	}
}
