package main

import (
	"fmt"
	"os"
	"strconv"
)

func update(rule, left, this, right int) int {
	index := left<<2 | this<<1 | right
	return (rule >> index) & 1
}

func to_string(cells []int) string {
	result := ""
	for _, cell := range cells {
		if cell == 0 {
			result += "░" // U+2591 (light shade)
		} else {
			result += "█" // U+2588 (full block)
		}
	}
	return result
}

func main() {
	if len(os.Args) != 4 {
		fmt.Printf("Usage: %s cell_count generation_limit rule\n", os.Args[0])
		return
	}
	cell_count, _ := strconv.Atoi(os.Args[1])
	generation_limit, _ := strconv.Atoi(os.Args[2])
	rule, _ := strconv.Atoi(os.Args[3])

	cells := make([]int, cell_count)
	cells[cell_count/2] = 1

	for generations := 0; generations < generation_limit; generations++ {
		fmt.Println(to_string(cells))

		new_cells := make([]int, cell_count)
		for i := 0; i < cell_count; i++ {
			left := cells[(i+cell_count-1)%cell_count]
			this := cells[i]
			right := cells[(i+1)%cell_count]
			new_cells[i] = update(rule, left, this, right)
		}
		cells = new_cells
	}
}
