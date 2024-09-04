package main

import (
    "fmt"
    "math"
)

type stack[T any] struct {
    Push    func(T)
    Pop     func() T
    PopLast func(int) []T
    Length  func() int
}

func Stack[T any]() stack[T] {
    slice := make([]T, 0)
    return stack[T]{
        Push: func(i T) {
            slice = append(slice, i)
        },
        Pop: func() T {
            res := slice[len(slice)-1]
            slice = slice[:len(slice)-1]
            return res
        },
        PopLast: func(n int) []T {
            res := slice[len(slice)-n:]
            slice = slice[:len(slice)-n]
            return res
        },
        Length: func() int {
            return len(slice)
        },
    }
}

func E(n int) int {
    return n * 2
}

func O(n int) int {
    return n*2 - 1
}

func T(n int) int {
    return int(math.Ceil((-1 + math.Sqrt(float64(1+8*n))) / 2))
}

func Combine(l func(int) int, r func(int) int) func(int) int {
    return func(n int) int {
        return r(l(r(n)))
    }
}

func main() {
    funcs := Stack[func(int) int]()
    bracks := Stack[int]() // Stores the positions of ( in the string
    // Get the string from the user
    var s string
    var x int
    _, err := fmt.Scanln(&s, &x)
    if err != nil {
        return
    }
    s = "(" + s + ")"
    for _, c := range s {
        switch c {
        case 'E':
            funcs.Push(E)
        case 'O':
            funcs.Push(O)
        case 'T':
            funcs.Push(T)
        case '(':
            bracks.Push(funcs.Length())
        case ')':
            n := funcs.Length() - bracks.Pop() // How many functions to combine

            farr := funcs.PopLast(n)

            cmb := farr[0]
            for i := 1; i < n; i++ {
                f := farr[i]
                cmb = Combine(cmb, f)
            }

            funcs.Push(cmb)
        }
    }

    fmt.Println(funcs.Pop()(x))
}