# Hello World

## Seu primeiro programa em Go

Uma vez com o GO devidamente instalado, podemos partir para o nosso primeiro programa! Você pode criar um arquivo com o código abaixo:

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}
```

e executá-lo com o comando:

```bash
go run <seu_arquivo>
```

Se tudo estiver correto, você verá a mensagem `Hello, World!` impressa no terminal!!

Parabéns, você acabou de executar o seu primeiro programa em Go! (Embora talvez não tenha entendido ainda 😅)

> Curiosidade: A tradição do "rito de passagem" do hello world começou em 1978 com o livro 'The C Programming Language', livro escrito pelos criadores do C que introduzia a linguagem, desde então, surgiu a superstição.

## Entendendo o código linha a linha

### Linha: `package main`

Todos os programas GO começam com a declaração de um pacote. Pacotes são a forma de modularização da linguagem (parecido com os packages de java, aos familiarizados).<br>
<br>
Não vamos nos aprofundar em pacotes aqui, mas eles são centrais para a organização semântica, gerenciamento de dependências e visibilidade (public, protected e private para os que já conhecem, aqui é diferente, não existem essas keywords, tipos e valores começando com letras maiúsculas são assumidos públicos e são acessíveis fora do pacote, aqueles começando com letras minúsculas são assumidos privados e são acessíveis por todos os arquivos do pacote, mas não fora dele).

#### Linha: `import "fmt"`

Aqui estamos apenas importando o pacote `fmt`, apesar do nome esquisito, fmt significa apenas format (a comunidade golang gosta de abreviações), é o pacote que tem funções de entrada e saída formatadas.<br>
Entre essas funções está o `Println`, que usamos para imprimir a mensagem no terminal, ele seria o equivalente ao `print` do python ou o `System.out.println` do java, etc.

#### Linha: `func main() {`

Para a surpresa de ninguém, essa linha declara uma função chamada main, que o GO usa como entrypoint do binário.<br>

#### Linha: `fmt.Println("Hello, World!")`

Alguns detalhes aqui:<br>

- Não estamos chamando um método de um objeto, mas chamando uma função definida em um pacote, por isso o formato `pacote.Função()`, note o nome começa com letra maiúscula, indicando que é pública.<br>
- Em GO não é possível fazer imports parciais, ou seja, não existe o conceito de importar apenas um método ou variável específica de um pacote, você sempre importa o pacote inteiro e usa o formato `pacote.Função()`.<br>
- Strings são sempre entre aspas duplas, aspas simples são usadas para caracteres.<br>
- As linhas são terminadas em `;`, mas elas são opcionais 🥳, normalmente só são colocadas quando existem múltiplas instruções na mesma linha.<br>

#### Linha: `}`

Supreendentemente, essa linha fecha o bloco da função main.
