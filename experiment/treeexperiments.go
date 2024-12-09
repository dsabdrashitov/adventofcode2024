package main

import (
	"fmt"
	"math/rand/v2"

	"github.com/dsabdrashitov/adventofcode2024/pkg/splaytree"
	"golang.org/x/exp/constraints"
)

func DoTreeExperiments() {
	tree := splaytree.New[int, string]()
	m := make(map[int]string)
	// m[0] = "!!!"
	fmt.Println("Test small")
	for range 100000 {
		k := rand.IntN(20)
		if rand.Float64() < 0.5 {
			v := fmt.Sprintf("value@%v", k)
			tree.Set(k, v)
			m[k] = v
		} else {
			tree.Delete(k)
			delete(m, k)
		}
		if !same(tree, m) {
			fmt.Println("Different:")
			fmt.Println(tree.Keys())
			fmt.Println(m)
		}
	}
	fmt.Println("Test large")
	for range 100000 {
		k := rand.IntN(100000000)
		if rand.Float64() < 0.5 {
			v := fmt.Sprintf("value@%v", k)
			tree.Set(k, v)
			m[k] = v
		} else {
			tree.Delete(k)
			delete(m, k)
		}
	}
	if !same(tree, m) {
		fmt.Println("Different:")
		fmt.Println(tree.Keys())
		fmt.Println(m)
	}

}

func same[K constraints.Ordered, V comparable, A any](tree *splaytree.SplayTree[K, V, A], m map[K]V) bool {
	for _, k := range tree.Keys() {
		v1, _ := tree.Get(k)
		v2, ok := m[k]
		if !ok || v1 != v2 {
			return false
		}
	}
	for k, v2 := range m {
		v1, ok := tree.Get(k)
		if !ok || v1 != v2 {
			return false
		}
	}
	return true
}
