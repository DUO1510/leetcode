package main

import "fmt"

func main() {
	fmt.Println(diStringMatch("IDID"))
}

// I perm[i] < perm[i+1]
// D perm[i] > perm[i+1]

func diStringMatch(s string) []int {
	res := make([]int, len(s)+1)
	ls := 0
	hs := len(s)
	for i, v := range s {
		if v == 'I' {
			res[i] = ls
			ls++
		} else {
			res[i] = hs
			hs--
		}
	}
	res[len(s)] = ls
	return res
}
