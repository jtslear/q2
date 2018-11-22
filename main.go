package main

import "fmt"

func theQuestion(i []int) []int {
	r := []int{}
	for a := 0; a < len(i); a++ {
		b := 0
		c := 1
		for b = 0; b < len(i); b++ {
			if a == b {
				continue
			} else {
				c = c * i[b]
			}
		}
		b = 0
		r = append(r, c)
	}
	return r
}

func main() {
	fmt.Println("vim-go")
}
