package main

import (
    "encoding/json"
    "fmt"
    "log"
    "os"
)


const (
    Reset  = "\033[0m"
    Red    = "\033[31m"
    Green  = "\033[32m"
    Yellow = "\033[33m"
    Blue   = "\033[34m"
    Cyan   = "\033[36m"
    Bold   = "\033[1m"
)

type Faturamento struct {
    Dia   int     `json:"dia"`
    Valor float64 `json:"valor"`
}


type FaturamentoStrategy interface {
    Calcular(faturamentos []Faturamento)
}


type MinMaxStrategy struct{}

func (m *MinMaxStrategy) Calcular(faturamentos []Faturamento) (min, max float64) {
    log.Println("Iniciando cálculo do menor e maior valor de faturamento...")
    min, max = faturamentos[0].Valor, faturamentos[0].Valor
    for _, f := range faturamentos {
        log.Printf("Analisando dia %d com valor de faturamento R$%.2f\n", f.Dia, f.Valor)
        if f.Valor > 0 {
            if f.Valor < min {
                log.Printf("Novo menor valor encontrado: R$%.2f no dia %d\n", f.Valor, f.Dia)
                min = f.Valor
            }
            if f.Valor > max {
                log.Printf("Novo maior valor encontrado: R$%.2f no dia %d\n", f.Valor, f.Dia)
                max = f.Valor
            }
        }
    }
    log.Println("Cálculo do menor e maior valor concluído.")
    return
}


type MediaStrategy struct{}

func (m *MediaStrategy) Calcular(faturamentos []Faturamento) (diasAcimaMedia int, media float64) {
    log.Println("Iniciando cálculo da média mensal e dias acima da média...")
    var soma float64
    var diasComFaturamento int

    for _, f := range faturamentos {
        if f.Valor > 0 {
            soma += f.Valor
            diasComFaturamento++
        }
    }

    media = soma / float64(diasComFaturamento)
    log.Printf("Média mensal calculada: R$%.2f\n", media)

    for _, f := range faturamentos {
        if f.Valor > media {
            log.Printf("Dia %d com faturamento R$%.2f acima da média.\n", f.Dia, f.Valor)
            diasAcimaMedia++
        }
    }
    log.Println("Cálculo da média e dias acima da média concluído.")
    return
}

func lerFaturamento(filePath string) ([]Faturamento, error) {
    log.Printf("Lendo arquivo de faturamento: %s\n", filePath)
    var faturamentos []Faturamento
    file, err := os.ReadFile(filePath)
    if err != nil {
        log.Printf("Erro ao ler o arquivo: %v\n", err)
        return nil, err
    }
    err = json.Unmarshal(file, &faturamentos)
    if err != nil {
        log.Printf("Erro ao fazer parse do JSON: %v\n", err)
        return nil, err
    }
    log.Println("Arquivo de faturamento lido e parseado com sucesso.")
    return faturamentos, nil
}

func main() {
    faturamentos, err := lerFaturamento("./faturamento.json")
    if err != nil {
        log.Println("Erro ao carregar o arquivo de faturamento. Encerrando execução.")
        return
    }

    
    minMax := &MinMaxStrategy{}
    min, max := minMax.Calcular(faturamentos)

    
    mediaStrategy := &MediaStrategy{}
    diasAcimaMedia, media := mediaStrategy.Calcular(faturamentos)

    
    fmt.Println(Bold + Blue + "\n======== Resultado do Faturamento ========" + Reset)
    fmt.Printf("%sMenor valor de faturamento:%s %sR$%.2f%s\n", Bold, Reset, Green, min, Reset)
    fmt.Printf("%sMaior valor de faturamento:%s %sR$%.2f%s\n", Bold, Reset, Green, max, Reset)
    fmt.Printf("%sMédia mensal:%s %sR$%.2f%s\n", Bold, Reset, Yellow, media, Reset)
    fmt.Printf("%sDias com faturamento acima da média:%s %s%d%s\n", Bold, Reset, Cyan, diasAcimaMedia, Reset)
    fmt.Println(Bold + Blue + "=========================================" + Reset)
}
