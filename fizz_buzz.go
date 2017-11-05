package main

import "fmt"

func fizzbuzz(number int) string {
    var msg string = ""
    if number % 3 == 0 && number % 5 == 0 {
        msg = "Fizz Buzz"
    } else if number % 3 == 0 {
        msg = "Fizz"
    } else if number % 5 == 0 {
        msg = "Buzz"
    }
    return msg
}

func main() {
    for i := 0; i < 100; i++ {
        fmt.Println(fizzbuzz(i))
    }
}
