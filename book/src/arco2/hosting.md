# Virando o jogo

Até então aprendemos uma noção sobre como a comunicação funciona, e fizemos nossas primeiras requisições para outros servidores. Chegou a hora de começar a ser o servidor!

## Idéia geral

Assim como no caso do `gurl` que fizemos, o GO vai tratar de todos os detalhes de implementação para nós, o que vamos precisar fazer é basicamente registrar um *callback* para lidar com a requisição, ou seja, uma *função* que vai ser chamada todas as vezes que uma nova requisição chegar.<br>

O código ficaria assim:


```go
package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", funcaoHandler) // Aqui dizemos que para todas as requisições que chegarem, funcaoHandler deve ser executada
	// o primeiro parâmetro é o padrão de match, ou seja, a URI que a request deve ter para cair nesse handler, "/" significa todas

	err := http.ListenAndServe("localhost:8000", nil) // Aqui é onde explicitamos inciamos o servidor, indicando em que porta ele deve inciar, o segundo parâmetro foge um pouco do escopo, mas seria um router

	// O nosso stub de tratamento de erro
	if err != nil {
		fmt.Println("erro inciando o servidor: ", err)
		os.Exit(1)
	}
}

// Essa é a função que será executada em todas as requisições
func funcaoHandler(
	writer http.ResponseWriter, // Aqui é onde escrevemos a resposta, é uma daquelas "streams"
	request *http.Request, // Esse é o objeto da request, existem muitas informações aqui, mas vamos usar poucas
) {
	helloEmBytes := []byte("Hello!\n")
	writer.Write(helloEmBytes) // Por hora escrevemos um 'Hello!'
}
```
<br>

Tente alterar o código e ver o que acontece, e tente entender porquê, por exemplo, se você remover a linha do listenAndServe, o que vai acontecer? Tente entender o que aconteceu<br>
Tente alterar também o código de `funcaoHandler`.

### __PRONTO__ já temos um servidor rodando e recebendo requisições! Você pode testar com o `curl` ou com o nosso `gurl` ou ainda abrir o navegador e verá a mensagem sendo impressa, lembrando que para acessar o servidor você usa localhost:<porta>

Agora os próximos passos vão ser adicionar comportamento no handler, os próximos exemplos vão usar esse server como base, é bom que você entenda todas as linhas presentes
