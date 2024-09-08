package main

import "fmt"

func inverterString(input string) string {
    runas := []rune(input)
    for i, j := 0, len(runas)-1; i < j; i, j = i+1, j-1 {
        runas[i], runas[j] = runas[j], runas[i]
    }
    return string(runas)
}

func main() {
    var input string
    fmt.Println("Informe uma string:")
    fmt.Scan(&input)

    fmt.Printf("String invertida: %s\n", inverterString(input))
}
