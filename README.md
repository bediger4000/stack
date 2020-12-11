# Make a stack from a heap

This is _Daily Coding Problem: Problem #438 [Easy]_

This problem was asked by Amazon.

Implement a stack API using only a heap. A stack implements the
following methods:

* push(item), which adds an element to the stack
* pop(), which removes and returns the most recently added element (or throws
an error if there is nothing on the stack)

Recall that a heap has the following operations:

* push(item), which adds a new key to the heap
* pop(), which removes and returns the max value of the heap

## Build and Run

```sh
$ go build stack.go
$ ./stack + 0 + 1 + 2 - - - - + 3
Popped 2, stack has 2 items
Popped 1, stack has 1 items
Popped 0, stack has 0 items
no values on stack
Left items on stack:
3
2020/09/13 13:26:05 no values on stack
```

The program `stack` reads arguments from command line.
A "+" followed by a seperate number representation
pushes that number on the stack.
A "-" pops the top off the stack and prints the value.
Note that the "+" and "-" are not concatenated,
there's whitespace between all command line arguments.

## My solution's algorithm

Keep track of an integer, as well as a heap.

Each time there's a stack push, increment the integer.
Use that integer as the key to the heap ordering.
That integer becomes an ordinal for the conceptual heap.
Popping the conceptual stack becomes removing the largest ordinal-valued
item from the heap.

Stack pop is the same as heap pop,
except that the integer is decremented.

## Analysis

This is an odd problem to solve.
A heap data structure with associated code is substantially slower
than the ususal array or linked list that one builds a stack from.
Maybe that's the point: to pose a problem that the candidate is
unlikely to ever have implemented,
even in practice for coding interviews.

The downside to the oddity is that it requires some programming
insight to use the stack push's index as the heap key.

The coding (after having the insight) is only marginally harder
than writing code for an array or linked-list based stack.
This is a question for a really sharp entry-level,
or a mid-level candidate.
The use of a heap will throw off some candidates,
as that "frames" the problem in an unusal way.
Sure, the coding is easy,
but getting to the coding requires a little of that weird
"use one thing as another" that a lot of algorithms require.
