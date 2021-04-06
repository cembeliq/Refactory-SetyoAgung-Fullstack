package main

import(
	"fmt"
)

func main() {
	a := makeRange(0, 1000)

	fmt.Println("==========EVEN=============")
	fmt.Println(even(a))
	fmt.Println()
	fmt.Println("==========ODD=============")
	fmt.Println(odd(a))
	fmt.Println()
	fmt.Println("==========MULTIPLY BY 5=============")
	fmt.Println(multiplyBy5(a))
	fmt.Println()
	fmt.Println("==========PRIME NUMBERS=============")
	fmt.Println(primeNumber(a))
	fmt.Println()
	fmt.Println("==========PRIME NUMBERS LESS THAN 100=============")
	a = makeRange(0, 100)
	fmt.Println(primeNumber(a))
}

func even(nums []int) []int {
	even := []int{}
	for i, num := range nums {
		if num % 2 == 0 {
			even = append(even, i)
		}
	}
	return even
}

func odd(nums []int) []int {
	odd := []int{}
	for i, num := range nums {
		if num % 2 == 1 {
			odd = append(odd, i)
		}
	}
	return odd
}

func multiplyBy5(nums []int) []int {
	multiplyBy5 := []int{}
	for i, num := range nums {
		if num % 5 == 0 {
			multiplyBy5 = append(multiplyBy5, i)
		}
	}
	return multiplyBy5
}

func primeNumber(nums []int) []int {
	primeNumber := []int{}
	for i, num := range nums {
		if num == 0 || num == 1 {
			continue
		} 
		if num % 2 == 0 && num != 2 {
			continue
		} else if num % 3 == 0 && num != 3{
			continue
		} else if num % 5 == 0 && num != 5{
			continue
		} else {
			primeNumber = append(primeNumber, i)
		}
	}
	return primeNumber
}

func makeRange(min, max int) []int {
    a := make([]int, max-min+1)
    for i := range a {
        a[i] = min + i
    }
    return a
}
