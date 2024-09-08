package main

import "fmt"

const INDICE = 13

func calcularSoma() int {
    soma, k := 0, 0
    for k < INDICE {
        k++
        soma += k
    }
    return soma
}

func main() {
    soma := calcularSoma()
    fmt.Printf("O valor final de SOMA é: %d\n", soma)
}
