# Assignment 02 - Functions & Packages

**Goal:** Create a package and a command-line tool to sort input provided by the user.
**Inputs:** Number (integer or float) array, string array.
**Outputs:** Sorted result based on the provided input type.

```bash
$ go run sorter.go -t int 5 2 10 1
# Output: [1 2 5 10]

$ go run sorter.go -t string apple orange banana
# Output: [apple banana orange]

$ go run sorter.go -t float 10.2 2.54 1.999 3.2991
# Output: [1.999 2.54 3.2991 10.2]
```