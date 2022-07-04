package main

import (
	"fmt"
	"github.com/willnotlazy/leetcode/questions"
)

func main() {
	target := []int{1, 2, 3, 4, 5, 10}
	fmt.Println(questions.SearchInsert(target, 2))
}
