package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 4, 5, 5}
	uniq := []int{}

	for _, v1 := range nums {
		found := false
		for _, v2 := range uniq {
			if v1 == v2 {
				found = true
				break

			}
		}
		if !found {
			uniq = append(uniq, v1)
		}
	}
	fmt.Println(uniq)
}