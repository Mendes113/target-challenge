package main

import "fmt"


type FibonacciChecker struct{}


func NewFibonacciChecker() *FibonacciChecker {
    return &FibonacciChecker{}
}


func (fc *FibonacciChecker) IsFibonacci(num int) bool {
    a, b := 0, 1
    for b <= num {
        if b == num {
            return true
        }
        a, b = b, a+b
    }
    return false
}

func main() {
    var num int
    fmt.Println("Informe um número:")
    fmt.Scan(&num)

    checker := NewFibonacciChecker()
    if checker.IsFibonacci(num) {
        fmt.Printf("O número %d pertence à sequência de Fibonacci.\n", num)
    } else {
        fmt.Printf("O número %d não pertence à sequência de Fibonacci.\n", num)
    }
}
