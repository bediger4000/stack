package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"stack/heap"
	"strconv"
)

func main() {
	stk := &stack{}
	for i := 0; i < len(os.Args[1:]); i++ {
		str := os.Args[i+1]
		if str == "+" {
			n, err := strconv.Atoi(os.Args[i+2])
			if err != nil {
				log.Fatal(err)
			}
			stk.push(n)
			i++
			continue
		}
		if str == "-" {
			n, err := stk.pop()
			if err != nil {
				fmt.Printf("%v\n", err)
				continue
			}
			fmt.Printf("Popped %d, stack has %d items\n", n, len(stk.heap))
		}
	}

	if stk.count > 0 {
		fmt.Println("Left items on stack:")
		for n, err := stk.pop(); true; n, err = stk.pop() {
			if err != nil {
				log.Print(err)
				break
			}
			fmt.Printf("%d\n", n)
		}
	}
}

type stack struct {
	heap  heap.Heap
	count int64
}

func (s *stack) push(value int) {
	s.count++
	s.heap = s.heap.Insert(&StackNode{Serial: s.count, Data: value})
}

func (s *stack) pop() (int, error) {
	if s.count < 1 {
		return 0, errors.New("no values on stack")
	}
	var n heap.Node
	s.heap, n = s.heap.Delete()
	s.count--
	sn := n.(*StackNode)
	return sn.Data, nil
}

type StackNode struct {
	Serial int64
	Data   int
}

func (n *StackNode) Value() int64 {
	return n.Serial
}

func (n *StackNode) IsNil() bool {
	return n == nil
}

func (n *StackNode) String() string {
	return fmt.Sprintf("%d", n.Data)
}
