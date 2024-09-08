package main

import "fmt"

func calcularPercentuais(faturamento map[string]float64) map[string]float64 {
    total := 0.0
    for _, valor := range faturamento {
        total += valor
    }

    percentuais := make(map[string]float64)
    for estado, valor := range faturamento {
        percentuais[estado] = (valor / total) * 100
    }

    return percentuais
}

func main() {
    faturamentoEstados := map[string]float64{
        "SP":     67836.43,
        "RJ":     36678.66,
        "MG":     29229.88,
        "ES":     27165.48,
        "Outros": 19849.53,
    }

    percentuais := calcularPercentuais(faturamentoEstados)

    fmt.Println("Percentual de representação por estado:")
    for estado, percentual := range percentuais {
        fmt.Printf("%s: %.2f%%\n", estado, percentual)
    }
}
