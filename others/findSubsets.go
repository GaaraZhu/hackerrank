package main

import (
	"fmt"
	"strings"
)

/*
 * Given an array of a list of unique numbers, try to find all its subsets
 */
func main() {
	ns := []int32{4, 6, 9, 15}
	ss := findSubsetsRecursively(ns)
	for _, s := range ss {
		fmt.Println(strings.Join(intArrayToStringArray(s)[:], ","))
	}
}

func intArrayToStringArray(ia []int32) []string {
	var sa []string
	for _, n := range ia {
		sa = append(sa, fmt.Sprint(n))
	}
	return sa
}

// solution 1
func findSubsets(ns []int32) [][]int32 {
	var subsets [][]int32
	for _, n := range ns {
		for _, s := range subsets {
			subsets = append(subsets, append(s, n))
		}
		subsets = append(subsets, []int32{n})
	}

	return subsets
}

// solution 2
func findSubsetsRecursively(ns []int32) [][]int32 {
	var lastElement = ns[len(ns)-1]
	if len(ns) == 1 {
		return [][]int32{[]int32{lastElement}}
	}

	var ssOfPrevious = findSubsetsRecursively(ns[0 : len(ns)-1])
	var subsets = ssOfPrevious
	for _, s := range ssOfPrevious {
		subsets = append(subsets, append(s, lastElement))
	}
	subsets = append(subsets, []int32{lastElement})

	return subsets
}
