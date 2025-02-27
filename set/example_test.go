package set

import "fmt"

func ExampleSet_Seq() {
	s := New(1, 2, 3, 4)
	for v := range s.Seq() {
		fmt.Println(v)
	}

	// Output:
	// 1
	// 2
	// 3
	// 4
}
