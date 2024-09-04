package main
import "fmt"

func Fibonacci(n int) []int {
    nums := make([]int, 0, n)
    nums = append(nums, 1)
    nums = append(nums, 2)

    for i := 2; nums[i - 1] < n; i++ {
        nums = append(nums, nums[i-1] + nums[i-2])
    }

    return nums
}

func Zeckendorf(n int) []int {
    fib := Fibonacci(n)
    zeck := make([]int, 0, n)

    i := len(fib) - 1
    for n > 0 {
        if fib[i] <= n {
            zeck = append(zeck, fib[i])
            n -= fib[i]
        }
        i--
    }

    return zeck
}

func main() {
    var n int
    fmt.Scan(&n)

    zeck := Zeckendorf(n)
    for _, v := range zeck {
        fmt.Print(v, " ")
    }
}