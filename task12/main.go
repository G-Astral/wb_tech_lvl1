// Имеется последовательность строк
// (cat, cat, dog, cat, tree)
// создать для нее собственное множество

package main

import (
	"fmt"
	"sort"
)

func main() {
	var resArr []string

	startArr := []string{"cat", "cat", "dog", "cat", "tree"}
	// startArr := []string{
	// 	"cat", "dog", "tree", "bird", "lion", "wolf", "tiger", "bear", "mouse", "fox",
	// 	"cat", "dog", "lion", "snake", "eagle", "fish", "whale", "shark", "owl", "bee",
	// 	"cat", "elephant", "lion", "wolf", "monkey", "horse", "rabbit", "fox", "panda", "koala",
	// 	"dog", "tree", "bird", "snake", "turtle", "shark", "crocodile", "penguin", "bee", "whale",
	// 	"wolf", "cat", "dog", "lion", "fox", "horse", "owl", "tiger", "bear", "mouse",
	// }

	uniqueMap := make(map[string]struct{})

	for _, v := range startArr {
		uniqueMap[v] = struct{}{}
	}

	for k := range uniqueMap {
		resArr = append(resArr, k)
	}

	sort.Strings(resArr)
	fmt.Println(resArr)
}
