# Integrando com um serviço externo

Por hora o nosso servidor apenas retorna um singelo `Hello!`, vamos estender ele para retornar a resposta de uma das três APIs que vimos quando estávamos construindo o client a partir da URI.<br>

A idéia vai ser:
- Caso a URI seja /chuck: retornar a piada aleatória
- Caso a URI seja /pokemon: retornar um pokemon aleatório (note que a API dos pokemons não tem essa funcionalidade)
- Caso a URI seja /rickandmorty: retornar um personagem do rick and morty aleatório (mesmo caso do de cima)<br>

Antes de continuarmos, vamos isolar a lógica de fazer a requisição em um arquivo separado, esse é o mesmo código do client que estávamos usando encapsulado numa função, você pode chamar do que quiser:

```go
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// Esse é o mesmo código do GURL que fizemos encapsulado numa função retornando os bytes ao invés de imprimir eles
// copiado e colado de arco1/gurl.go

func enviarRequisicao(url string) []byte {
	resposta, err := http.Get(url) // A boa notícia é que GO já tem toda a parte complicada pronta, então podemos só usar!!
	if err != nil {
		fmt.Println("nao foi possível se comunicar com o servidor: ", err)
		os.Exit(1)
	}

	// Só fazemos ler a reposta e imprimir na tela
	bytesDaResposta := lerResposta(resposta)
	return bytesDaResposta
}

func lerResposta(resp *http.Response) []byte {
	body, err := io.ReadAll(resp.Body) // Aqui lemos a resposta do servidor para um array de bytes
	resp.Body.Close()                  // Precisamos sempre fechar as "streams" em GO (isso seria discussão para um outro workshop)

	if err != nil {
		fmt.Println("Não foi possível ler a resposta: ", err)
		os.Exit(1)
	}

	return body // o que é uma string se não um array de bytes?
}
```

### Integrando

Você talvez já deva imaginar como ficariam os handlers para cada uma das APIs agora. Ficaria algo assim:

```go
// server v1
package main

import (
	"fmt"
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

	handleAPI(writer, url)
}

func handlerPokemon(
	writer http.ResponseWriter,
	request *http.Request,
) {
	// Podemos gerar um número aleatório para buscar a cada request
	randomNumber := rand.Intn(100) + 1
	url := "hthttps://pokeapi.co/api/v2/pokemon/" + strconv.Itoa(randomNumber) // converte o num pra string e concatena

	handleAPI(writer, url)
}

func handlerRickAndMorty(
	writer http.ResponseWriter,
	request *http.Request,
) {
	// Podemos gerar um número aleatório para buscar a cada request
	randomNumber := rand.Intn(100) + 1
	url := "https://rickandmortyapi.com/api/character/" + string(randomNumber)

	handleAPI(writer, url)
}

// Chamado por todos
func handleAPI(writer http.ResponseWriter, url string) {
	resposta := enviarRequisicao(url)
	writer.Write(resposta)
}

// Essa é a função que será chamada caso a URI não seja nenhuma das outras
func handlerDefault(
	writer http.ResponseWriter,
	request *http.Request,
) {
	helloEmBytes := []byte("Oii, para acessar alguma API, use /chuck, /pokemon ou /rickandmorty!\n")
	writer.Write(helloEmBytes)
}

```

> Se você preferir separar o arquivo do client, você precisará passar esse arquivo como parâmetro para o go run também! Caso contrário ocorrerá erro

Simples, né? Essa é a nossa V1, mas ela tem um problema grave! Você consegue imaginar o que vai acontecer se ocorrer algum erro ao enviar a requisição para alguma API? Tente simular colocando a string errada e veja o que vai acontecer rsrs


### Tratando erros

O problema é que o client está finalizando o processo ao encontrar o menor erro, isso não é muito interessante...,

Para isso, vamos seguir o padrão idiomático do GO, a função enviarRequisicao passa a retornar duas coisas, um *[]byte* e um *err* indicando um possível problema que possa ter ocorrido.<br>
O handlers agora fariam algo parecido como:
```go
// server v2
handleAPI(writer http.ResponseWriter, url string) {
	resposta, err := enviarRequisicao(url)

	if err != nil {
		writer.Write([]byte(err.Error()))
		return
	}

	writer.Write(resposta)
}
```
(os arquivos em arco2/server_v2 tem as implementações atualizadas)

### Lendo a resposta e montando payload

Podemos retornar por exemplo o tamanho do resultado em um outro campo.

Vamos retornar um JSON com dois campos, { tamanho: number; resultado: any }, mas como podemos fazer isso?<br>

Para o tamanho, podemos apenas chamar o length<br>
Para retornar o dado estruturado em json, fazemos um Marshal (ou serialização) de uma struct anonima combinando essas duas variáveis, o handleAPI ficaria assim:

```go
// server v3

// Chamado por todos
func handleAPI(writer http.ResponseWriter, url string) {
	respostaDaApi, err := enviarRequisicao(url)
	if err != nil {
		writer.Write([]byte(err.Error()))
	}

	// Aqui transformamos os bytes retornados em objeto estruturado para serializarmos depois
	var respostaDaApiEmJson any
	json.Unmarshal(respostaDaApi, &respostaDaApiEmJson) // o unmarshal vai montar uma struct com os campos do json

	// Aqui tem um conceito novo, structs anonimas, parecidas com os types do typescript
	// IMPORTANTE: os campos precisam começar com letra maiúscula, caso contrário
	// seriam tratados como campos privados e não seriam serializados!
	resposta := struct {
		Tamanho  int
		Resposta any // isso significa que a tipagem desse campo é desconhecida
	}{
		Tamanho:  len(respostaDaApi),
		Resposta: respostaDaApiEmJson,
	}

	respostaJson, err := json.Marshal(resposta)
	if err != nil {
		writer.Write([]byte(err.Error()))
		return
	}

	writer.Write(respostaJson)

```

Com isso nós temos um pequeno servidor com diversas coisas já funcionando! Temos roteamento de requests, integramos com serviços externos e trabalhamos serialização/desserialização.

Se você quiser testar como seu server lida com várias requests ao mesmo tempo, pode usar o script em `arco2/requests.sh` passando a quantidade de requests como parâmetro
