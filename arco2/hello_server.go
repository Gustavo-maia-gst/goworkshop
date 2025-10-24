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
