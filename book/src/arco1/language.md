# Go — Mini Guia Rápido

## Variáveis
- Declaração explícita:
```go
var x int = 10
````

- Declaração com inferência de tipo:
```go
var x = 10 // tipo de x inferido como int
````

* Declaração curta (mais comum):

```go
y := 42 // y declarado como int e inicializado com 42
```

* Múltiplas variáveis:

```go
a, b := 1, "texto" // a é int = 1, b é string = "texto"
```

* Constantes:

```go
const pi = 3.14 // constantes conhecidas em compile time
```

---

## Tipos básicos

* Tipos primitivos: `int`, `float64`, `string`, `bool`, `rune`, `byte`
* Arrays e slices:

```go
array := [3]int{1,2,3}   							// array com tamanho estático 3
slice_inicializado := []int{1,2,3}    // slice dinâmico (listas do python, ou arrays do javascript)
slice_declarado := make([]int, 5) 		// slice pre-alocado com 5 elementos, mas vazio
var slice_nao_inicializado []int 			// slice vazio, referência para nil
```

* Maps:

```go
m := map[string]int{"a":1, "b":2}
```

* nil: É o null ou None do GO, usado para ponteiro sem valor

> 💡 GO é mais estrito com a possibilidade de um valor ser nulo do que outras linguagens, por exemplo, uma variável que aponta pra struct, string, int, etc NUNCA pode ser nil, apenas tipos de referência como ponteiros, mapas e slices podem ser nil.

---

## Funções

* Declaração simples:

```go
func soma(a, b int) int {
    return a + b
}
```

* Funções em GO são **valores de primeira classe**, podem ser atribuídas a variáveis:

```go
f := func(x int) int { return x*2 }
fmt.Println(f(3))  // 6
```

* E passadas como parâmetros
```go
func aplica(f func(int) int, x int) int {
    return f(x)
}
```

---

## Structs

* Definem **tipos compostos**:

```go
type Pessoa struct {
    Nome string
    Idade int
}
```

* Instanciando e acessando campos:

```go
p := Pessoa{Nome:"Gustavo", Idade:19}
fmt.Println(p.Nome)  // Gustavo
```

* Métodos

```go
func (p Pessoa) Saudacao() string {
    return "Olá, " + p.Nome
}

func (p Pessoa) SaudacaoPara(outro string) string { // note que o nome do argumento vem sempre antes do tipo
    return "Olá, " + outro + ", eu sou " + p.Nome
}

fmt.Println(p.Saudacao())  // Olá, Gustavo
fmt.Println(p.SaudacaoPara("Ana"))  // Olá, Ana, eu sou Gustavo
```

> 💡 Embora chamados de métodos, eles são bem diferentes dos métodos de objetos em linguagens orientadas a objeto (GO não é orientado a objeto), você pode imaginar esses métodos apenas como sintax sugar para funções que atuam sobre uma struct específica. Aqui não existe overload de método ou os outros problemas trazidos por espalhar o código entre as definições dos tipos como acontece em linguagens OO tradicionais.

---

## Expressões e operadores

* Aritméticos: `+ - * / %`
* Comparação: `== != < <= > >=`
* Lógicos: `&& || !`

```go
if x > 0 && x < 10 { ... } // sem parenteses desnecessários
```

---

## Controle de fluxo

* `if`, `else`, `switch`, `for` (não existe `while`)
* Loop clássico:

```go
for i := 0; i < 5; i++ {
    fmt.Println(i)
}
```

* Loop estilo “while”:

```go
i := 0
for i < 5 {
    fmt.Println(i)
    i++
}
```

* loop estilo range do python
```go
for i, item := range lista {
	fmt.Println("Item " + string(item) + " na posição " + string(i))
}
```

---

## Interfaces e polimorfismo

* Define comportamento que structs podem implementar:

```go
type Saudavel interface {
    Saudacao() string
}

func cumprimenta(s Saudavel) {
    fmt.Println(s.Saudacao())
}
```

* Structs implementam interface **implicitamente** se tiverem os métodos exigidos

> 💡 As interfaces em GO são apenas contratos, não exigem declaração explícita, basta que a struct tenha os métodos com o mesmo nome e assinatura, isso torna a linguagem muito flexível e deixa tudo menos verboso, embora adicione um pouco de complexidade, são os trade-offs típicos de design de linguagens.
---
