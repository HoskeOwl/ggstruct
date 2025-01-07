# goCollectionsGeneric
Data structures has had made with generics:

* queue
* set
* stack
* trie

## Examples
### queue
```golang
package main

import (
	"fmt"

	"github.com/HoskeOwl/goCollectionsGeneric/queue"
)

func main() {
	q := queue.New[int]()
	if _, exists := q.Dequeue(); !exists {
		fmt.Println("Empty")
	}
	q.Enqueue(34)
	if val, exists := q.Peek(); exists {
		fmt.Println(val)
	}
	q.Enqueue(12)
	if val, exists := q.Dequeue(); exists {
		fmt.Println(val)
	}
}
```
### set
```golang
package main

import (
	"fmt"

	"github.com/HoskeOwl/goCollectionsGeneric/set"
)

func main() {
	s1 := set.New[int]()
	s1.Insert(1)
	s1.Insert(2)
	s1.Insert(3)
	s2 := set.New[int]()
	s2.Insert(3)
	s2.Insert(4)
	s2.Insert(5)
	fmt.Println(s1.Difference(s2))
	fmt.Println(s1.Union(s2))
	fmt.Println(s1.Intersection(s2))
}
```
### stack
```golang
package main

import (
	"fmt"

	"github.com/HoskeOwl/goCollectionsGeneric/stack"
)

func main() {
	st := stack.New[int]()
	if _, exists := st.Peek(); !exists {
		fmt.Println("Empty")
	}
	st.Push(1)
	if val, exists := st.Peek(); exists {
		fmt.Println(val)
	}
	st.Push(3)
	if val, exists := st.Pop(); exists {
		fmt.Println(val)
	}
}
```