# Summary (Sumário)
- <a href="#pt-br">PT-BR</a>
    - <a href="#iniciando">Iniciando</a>
    - <a href="#como-usar">Como Usar</a>
        - <a href="#ramusage">RamUsage</a>
        - <a href="#totalram">TotalRam</a>
        - <a href="#ssd">SSD</a>
- <a href="#en-us">EN-US</a>
    - <a href="#getting-started">Getting Started</a>
    - <a href="#how-to-use">How To Use</a>
        - <a href="#ramusage-1">RamUsage</a>
        - <a href="#totalram-1">TotalRam</a>
        - <a href="#ssd-1">SSD</a>

# PT-BR
Aqui você encontrará a documentação para esse pacote, escrita em Português.

## Iniciando
Para começar, simplesmente faça o download do pacote usando o seguinte comando:
```bash
go get github.com/End313234/squarecloud-golang-status
```

## Como Usar

### RamUsage
Retorna o atual consumo de ram.
Exemplo:
```go
usage, err := squarecloud.RamUsage("KB") // 'kb' também é válido
if err != nil {
    fmt.Println(err)
    return
}
fmt.Println(usage) // 23000
```

### TotalRam
Retorna o total de ram disponível para o uso.
Exemplo:
```go
total, err := squarecloud.TotalRam("MB") // 'mb' também é válido
if err != nil {
    fmt.Println(err)
    return
}
fmt.Println(total) // 300
```

### SSD
Retorna o uso atual do SSD.
Exemplo:
```go
usage, err := squarecloud.RamUsage("KB") // 'kb' também é válido
if err != nil {
    fmt.Println(err)
    return
}
fmt.Println(usage) // 20000
```

# EN-US
Here you'll find the documentation written in English for this package.

## Getting Started
To get started, simply download the package using the following command:
```bash
go get github.com/End313234/squarecloud-golang-status
```

## How To Use

### RamUsage
Returns the current ram usage.
Example:
```go
package main

import (
    "github.com/squarecl/squarecloud-golang-status"
    "fmt"
)

func main() {
    usage, err := squarecloud.RamUsage("KB") // 'kb' is also valid
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(usage) // 23000
}
```

### TotalRam
Returns the avaliable ram.
Example:
```go
total, err := squarecloud.TotalRam("MB") // 'mb' is also valid
if err != nil {
    fmt.Println(err)
    return
}
fmt.Println(total) // 300
```

### SSD
Returns the current SSD usage.
Example:
```go
usage, err := squarecloud.RamUsage("KB") // 'kb' is also valid
if err != nil {
    fmt.Println(err)
    return
}
fmt.Println(usage) // 20000
```