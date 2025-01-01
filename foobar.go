package main

import "fmt"

// rules:       1. do not print prime numbers
//              2. if divisible by 3 -> Foo
//              3. if divisible by 5 -> Bar
//              4. if divisible by both 3 and 5 -> FooBar

func isPrime(number int) bool {
        // only need to test the number up to its square root
        if number <= 1 {
                return false
        }

        cnt := 0
        for i := 2; i * i <= number; i++ {
                if number % i == 0 {
                        cnt++
                }
        }

        if cnt > 0 {
                return false
        } else {
                return true
        }
}

func main () {
        rules := map[int]string{}       // i implemented the rules in map data type to easily add new rule if needed to.
                                        // the divisors are keys, "Foo" or "Bar" are values
        rules[3] = "Foo"
        rules[5] = "Bar"

        // from email: "Please create an array/list of numbers from 1 to 100."
        numbers := make([]int, 100)
        for i := 0; i < 100; i++ {
                numbers[i] = i+1;
        }

        // then iterate in reverse order
        separator := ""
        for i := len(numbers)-1; i >= 0; i-- {

                // first test if it is a prime number
                if isPrime(numbers[i]) {
                        continue
                }

                output_string := ""
                for divisor := range rules { // divisor is the key
                        if numbers[i] % divisor == 0 {
                                output_string += rules[divisor]
                        }
                }

                if output_string != "" {
                        fmt.Print(separator, output_string)
                } else {
                        fmt.Print(separator, numbers[i])
                }
                separator = ", "
        }
}
