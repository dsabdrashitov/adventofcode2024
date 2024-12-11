package main

import (
	"fmt"
	"math/rand/v2"
	"time"

	"github.com/dsabdrashitov/adventofcode2024/pkg/splaytree"
	"golang.org/x/exp/constraints"
)

func DoLongTree() {
	dataSize := 1000000
	testsCount := 1000000
	tree := splaytree.New[int, string]()
	m := make(map[int]string)
	for k := range dataSize {
		v := fmt.Sprintf("value@%v", k)
		tree.Set(k, v)
		m[k] = v
	}
	var keys []int

	keys = make([]int, testsCount)
	for i := range testsCount {
		if i%2 == 0 {
			keys[i] = 0
		} else {
			keys[i] = dataSize - 1
		}
	}

	testMap(m, keys)
	testTree(tree, keys)

	keys = make([]int, testsCount)
	for i := range testsCount {
		keys[i] = rand.IntN(dataSize)
	}

	testMap(m, keys)
	testTree(tree, keys)

	if !same(tree, m) {
		fmt.Println("Different:")
		fmt.Println(tree.Keys())
		fmt.Println(m)
	}
}

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
	var time0 time.Time
	time0 = time.Now()
	for _, k := range tree.Keys() {
		v1, _ := tree.Get(k)
		v2, ok := m[k]
		if !ok || v1 != v2 {
			return false
		}
	}
	fmt.Println("check tree => map", time.Since(time0))
	time0 = time.Now()
	for k, v2 := range m {
		v1, ok := tree.Get(k)
		if !ok || v1 != v2 {
			return false
		}
	}
	fmt.Println("check map => tree", time.Since(time0))
	return true
}

func testTree[A any](tree *splaytree.SplayTree[int, string, A], keys []int) {
	time0 := time.Now()
	for _, k := range keys {
		if _, ok := tree.Get(k); !ok {
			panic("")
		}
	}
	fmt.Println(time.Since(time0))
}

func testMap(m map[int]string, keys []int) {
	time0 := time.Now()
	for _, k := range keys {
		if _, ok := m[k]; !ok {
			panic("")
		}
	}
	fmt.Println(time.Since(time0))
}
