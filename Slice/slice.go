package main

import "fmt"

const (
	val1 = (iota + 1) * 10
	val2
)

func main() {
	var (
		sliceCapacity uint16 = 5
	)

	nums := make([]int, 0, sliceCapacity)

	fmt.Println("After make — len:", len(nums), "cap:", cap(nums))
	fmt.Println("Content:", nums)

	values := []int{val1, val2}

	for _, v := range values {
		nums = append(nums, v)
	}

	fmt.Println("After loop — len:", len(nums), "cap:", cap(nums))
	fmt.Println("Content:", nums)

	free := cap(nums) - len(nums)
	switch free {
	case 0:
		fmt.Println("Full: no room left (len == cap)")
		fallthrough
	default:
		fmt.Println("Free slots:", free)
	}
}
