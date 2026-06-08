# Curso: GO Expert - Full Cycle <img src="https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/go/go-original.svg" width="50" height="50"/>
Repositório com arquivos desenvolvidos durante as aulas do curso GO Expert da escola Full Cycle
### Buildar/Gerar um executável do arquivo
- antes de tudo, salvar o arquivo
- comando: go build + nome_arquivo.go

### Executar o arquivo
- comando: go run + nome_arquivo.go

### Palavras Reservadas na Linguagem Go (Golang)
A linguagem Go possui **25 palavras reservadas**. Essas palavras são utilizadas pela própria linguagem para definir sua sintaxe e estrutura. Portanto, elas **não podem ser utilizadas como identificadores** (nomes de variáveis, funções, ou quaisquer outros elementos definidos pelo usuário).

#### Lista de Palavras Reservadas
|              |              |              |               |               |
|--------------|--------------|--------------|---------------|---------------|
| `break`      | `default`    | `func`       | `interface`   | `select`      |
| `case`       | `defer`      | `go`         | `map`         | `struct`      |
| `chan`       | `else`       | `goto`       | `package`     | `switch`      |
| `const`      | `fallthrough`| `if`         | `range`       | `type`        |
| `continue`   | `for`        | `import`     | `return`      | `var`         |
|              |              |              |               |               |

#### Observações

- Go é uma linguagem **case-sensitive**, ou seja, diferencia maiúsculas de minúsculas. Por exemplo, `for` é uma palavra reservada, mas `For` não é.
- Essas palavras são essenciais para a construção de programas em Go e devem ser usadas conforme sua finalidade específica.

#### Exemplos de Uso

Aqui está um exemplo simples de como algumas dessas palavras reservadas são usadas em um programa Go:

```go
package main

import "fmt"

func main() {
    for i := 0; i < 5; i++ {
        if i%2 == 0 {
            fmt.Println(i, "é par")
        } else {
            fmt.Println(i, "é ímpar")
        }
    }
}
