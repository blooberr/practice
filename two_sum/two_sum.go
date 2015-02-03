package main

import (
	"fmt"
)

func main() {
	input := []int{2, 7, 11, 15}
	TwoSum(9, input)
}

func TwoSum(target int, input []int) {
	m := make(map[int]int)
	for index, value := range input {
		m[value] = index
	}

	for key, _ := range m {
	    lookup := target-key
		val, ok := m[lookup]
		if !ok {
			continue
		} else {
			fmt.Printf("key is %d val is %d \n", key, m[key])
			fmt.Printf("other key is %d %d \n", lookup, val)
		}
    }
}
