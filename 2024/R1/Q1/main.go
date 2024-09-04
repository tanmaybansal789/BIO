package main
import "fmt"

func NumDigits(x int) int {
	if x == 0 {
		return 1
	}
	i := 0
	for x != 0 {
		x /= 10
		i++
	}
	return i
}

func IntPow(x, y int) int {
	result := 1
	for y > 0 {
		if y&1 == 1 {
			result *= x
		}
		y >>= 1
		x *= x
	}
	return result
}

func DigitsInRange(x int) int {
	return 9 * IntPow(10, x) * (x + 1)
}

func DigitsTo(x int) int {
	n := NumDigits(x)

	r := 0
	for i := 0; i < n - 1; i++ {
		r += DigitsInRange(i)
	}

	r += (x - IntPow(10, n - 1)) * n

	return r
}

func NthDigit(x, i int) int {
	return x / IntPow(10, NumDigits(x) - i - 1) % 10
}

func DigitAt(i int) int {
	dl := 0
	n := 0
	for dl < i {
		nr := DigitsInRange(n)
		if dl + nr >= i {
			break
		}
		dl += nr
		n++
	}

	i -= dl
	sr := IntPow(10, n)
	num := sr + (i - 1) / (n + 1)
	numi := (i - 1) % (n + 1)

	return NthDigit(num, numi)
}

func DigitFromAt(x, i int) int {
	return DigitAt(DigitsTo(x) + i)
}

func main() {
	var x, i int
	_, err := fmt.Scanln(&x, &i)
	if err != nil {
		return
	}

	fmt.Println(DigitFromAt(x, i))
}
