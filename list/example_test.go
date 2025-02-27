package list

import (
	"fmt"
	"iter"
)

func ExampleList_Seq() {
	lst := New(1, 2, 3, 4)
	for num := range lst.Seq() {
		fmt.Println(num)
	}
	// Output:
	// 1
	// 2
	// 3
	// 4
}

func ExampleList_ReversedSeq() {
	lst := New(1, 2, 3, 4)
	for num := range lst.ReversedSeq() {
		fmt.Println(num)
	}
	// Output:
	// 4
	// 3
	// 2
	// 1
}

func ExampleList_Seq2() {
	lst := New(1, 2, 3, 4)
	for i, num := range lst.Seq2() {
		fmt.Printf("%d: %d\n", i, num)
	}
	// Output:
	// 0: 1
	// 1: 2
	// 2: 3
	// 3: 4
}

func ExampleList_ReversedSeq2() {
	lst := New(1, 2, 3, 4)
	for i, num := range lst.ReversedSeq2() {
		fmt.Printf("%d: %d\n", i, num)
	}

	fmt.Println()

	// If you need indexes from 0 to 3
	reverseIdx := func(seq2 iter.Seq2[int, int]) iter.Seq2[int, int] {
		return func(yield func(int, int) bool) {
			i := 0
			for _, v := range seq2 {
				if !yield(i, v) {
					return
				}
				i++
			}
		}
	}

	for i, num := range reverseIdx(lst.ReversedSeq2()) {
		fmt.Printf("%d: %d\n", i, num)
	}
	// Output:
	// 3: 4
	// 2: 3
	// 1: 2
	// 0: 1
	//
	// 0: 4
	// 1: 3
	// 2: 2
	// 3: 1

}
