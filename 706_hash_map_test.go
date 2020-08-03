package leetcode_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Node struct {
	Key   int
	Value int
	Next  *Node
}

type MyHashMap struct {
	Arrays   []*Node
	HashFunc func(key int) int
	Size     uint32
}

/** Initialize your data structure here. */
func Constructor() MyHashMap {
	size := 100
	hm := MyHashMap{
		Arrays: make([]*Node, size),
		Size:   100,
		HashFunc: func(key int) int {
			return key % size
		},
	}

	return hm
}

func (hm *MyHashMap) NewNode(key int, value int) *Node {
	return &Node{
		Key:   key,
		Value: value,
	}
}

/** value will always be non-negative. */
func (hm *MyHashMap) Put(key int, value int) {
	var prev *Node
	index := hm.HashFunc(key)
	entry := hm.Arrays[index]
	for entry != nil && entry.Key != key {
		prev = entry
		entry = entry.Next
	}

	if entry == nil {
		if prev == nil {
			prev = hm.NewNode(key, value)
			hm.Arrays[index] = prev
		} else {
			entry = hm.NewNode(key, value)
			prev.Next = entry
		}
	} else {
		entry.Value = value
	}

}

/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (hm *MyHashMap) Get(key int) int {
	entry := hm.Arrays[hm.HashFunc(key)]
	for entry != nil {
		if entry.Key == key {
			return entry.Value
		}

		entry = entry.Next
	}

	return -1
}

/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (hm *MyHashMap) Remove(key int) {
	var prev *Node
	entry := hm.Arrays[hm.HashFunc(key)]
	for entry != nil && entry.Key != key {
		prev = entry
		entry = entry.Next
	}

	if entry == nil {
		return
	}

	if prev == nil {
		hm.Arrays[hm.HashFunc(key)] = entry.Next
	} else {
		prev.Next = entry.Next
	}

	return
}

func Test706HashTable(t *testing.T) {
	obj := Constructor()
	obj.Put(1, 2)
	obj.Put(2, 2)
	obj.Put(2, 1)
	obj.Put(3, 3)
	assert.Equal(t, 2, obj.Get(1))
	assert.Equal(t, 1, obj.Get(2))
	assert.Equal(t, 3, obj.Get(3))
	obj.Remove(2)
	assert.Equal(t, obj.Get(2), -1)
}
