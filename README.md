# ggstruct 
Data structures has had made with generics:

* list
* queue
* set
* stack
* trie

`list` and `set` have iterators - can be used with `range`.

# Installation && import

`go get github.com/HoskeOwl/ggstruct`

`import "github.com/HoskeOwl/ggstruct`


# Benefits
1) No pointer conversions generics used (generic used). No manual type-checking. All structures can hold any `comparable` type.
2) `list` and `set` have iterators. Can be used with `range`.
3) `list` operate values not upper level objects like nodes.
4) `trie` can operate with any string as a key (unicode-tolerance).

## Examples
### set
```golang
package main

import (
	"fmt"

	"github.com/HoskeOwl/ggstruct/set"
)

func main() {
	s1 := set.New[int]()
	s1.Insert(1)
	s1.Insert(2)
	s1.Insert(3)
    
	s2 := set.New(3,4,5)
    
    
	fmt.Println(s1.Difference(s2))
	fmt.Println(s1.Union(s2))
	fmt.Println(s1.Intersection(s2))
    for num := range s2.Seq(){
        fmt.Println(num)
    }
}
```
### list
```golang
package main

import (
	"fmt"

	"github.com/HoskeOwl/ggstruct/list"
)

func main() {
	s1 := list.New[int]()
	s1.Add(1)
	s1.Add(2)
	s1.Add(3)
    
	s2 := list.New(3,4,5)
    
    for num := range s2.Seq(){
        fmt.Println(num)
    }
    
	for i, num := range s2.Seq2(){
		fmt.Printf("%d: %d", i, num)
	}
}
```
### stack
```golang
package main

import (
	"fmt"

	"github.com/HoskeOwl/ggstruct/stack"
)

func main() {
	st := stack.New[int]()
	if _, exists := st.Top(); !exists {
		fmt.Println("Empty")
	}
	st.Push(1)
	if val, exists := st.Top(); exists {
		fmt.Println(val)
	}
	st.Push(3)
	if val, exists := st.Pop(); exists {
		fmt.Println(val)
	}
}
```
### queue
```golang
package main

import (
	"fmt"

	"github.com/HoskeOwl/ggstruct/queue"
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
}
```
