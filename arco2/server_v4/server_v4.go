package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
)

func main() {
	// Cadastramos os 4 patterns, ao receber uma requisição com /pokemon na URI, somente o handlerPokemon será chamado
	http.HandleFunc("/", handlerDefault)
	http.HandleFunc("/chuck", handlerChuck)
	http.HandleFunc("/pokemon", handlerPokemon)
	http.HandleFunc("/rickandmorty", handlerRickAndMorty)

	err := http.ListenAndServe("localhost:8000", nil)

	// Novamente o nosso stub de tratamento de erro
	if err != nil {
		fmt.Println("erro inciando o servidor: ", err)
		os.Exit(1)
	}
}

// Os 3 handlers fazem a mesma coisa, fazem a requisição para a API externa e retornam os resultados!
func handlerChuck(
	writer http.ResponseWriter,
	request *http.Request,
) {
	// Podemos gerar um número aleatório para buscar a cada request
	url := "https://api.chucknorris.io/jokes/random"

	handleAPI(writer, request, url)
}

func handlerPokemon(
	writer http.ResponseWriter,
	request *http.Request,
) {
	// Podemos gerar um número aleatório para buscar a cada request
	randomNumber := rand.Intn(100) + 1
	url := "https://pokeapi.co/api/v2/pokemon/" + strconv.Itoa(randomNumber) // converte o num pra string e concatena

	handleAPI(writer, request, url)
}

func handlerRickAndMorty(
	writer http.ResponseWriter,
	request *http.Request,
) {
	// Podemos gerar um número aleatório para buscar a cada request
	randomNumber := rand.Intn(100) + 1
	url := "https://rickandmortyapi.com/api/character/" + strconv.Itoa(randomNumber)

	handleAPI(writer, request, url)
}

// Chamado por todos
func handleAPI(writer http.ResponseWriter, request *http.Request, url string) {
	requestId := smallID(12) // gera um ID aleatório
	logger := log.New(os.Stdout, request.RequestURI, log.LstdFlags)
	logger.SetPrefix("[" + requestId + "] " + request.RequestURI + "\t")

	logger.Println("Requisição recebida")

	logger.Println("Enviando requisição para API externa")

	respostaDaApi, err := enviarRequisicao(url)
	if err != nil {
		writer.Write([]byte(err.Error()))
	}

	logger.Println("Resposta recebida")

	// Aqui transformamos os bytes retornados em objeto estruturado para serializarmos depois
	var respostaDaApiEmJson any
	json.Unmarshal(respostaDaApi, &respostaDaApiEmJson) // o unmarshal vai montar uma struct com os campos do json

	// Aqui tem um conceito novo, structs anonimas, parecidas com os types do typescript
	// IMPORTANTE: os campos precisam começar com letra maiúscula, caso contrário
	// seriam tratados como campos privados e não seriam serializados!
	resposta := struct {
		RequestId string
		Tamanho   int
		Resposta  any // isso significa que a tipagem desse campo é desconhecida
	}{
		RequestId: requestId,
		Tamanho:   len(respostaDaApi),
		Resposta:  respostaDaApiEmJson,
	}

	logger.Println("Montando resposta")

	respostaJson, err := json.Marshal(resposta)
	if err != nil {
		writer.Write([]byte(err.Error()))
		return
	}

	writer.Write(respostaJson)

	logger.Println("Resposta enviada")
}

// Essa é a função que será chamada caso a URI não seja nenhuma das outras
func handlerDefault(
	writer http.ResponseWriter,
	request *http.Request,
) {
	helloEmBytes := []byte("Oii, para acessar alguma API, use /chuck, /pokemon ou /rickandmorty!\n")
	writer.Write(helloEmBytes)
}

func smallID(n int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}
