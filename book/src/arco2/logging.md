# Introduzindo o conceito de logging

Uma coisa que talvez tenha te incomodado, é que não temos informação sobre o que está acontecendo no servidor... Não sabemos dizer onde está a maior demora no processamento e esse tipo de coisa.<br>

Esse problema é resolvido por meio de logging, isso é tão importante que o GO traz na sua stdlib um package para logging, que vamos usar para deixar o nosso servidor com uma cara um pouco mais profissional.

Para isso, importamos o package `log`, e loggamos informações relevantes com `log.Default().`. Uma coisa muito interessante do log, é que podemos criar objetos de logging com propriedades definidas, por exemplo, o snippet abaixo vai criar uma nova instância de logger e setar nela um prefixo, todas as chamadas subsequentes de log será prefixada com o texto passado!

```go
// server v4
logger := log.New(os.Stdout, request.RequestURI, log.LstdFlags)
logger.SetPrefix(request.RequestURI + "\t")
```

Adicione alguns logs ao seu código para rastrear o que está sendo executado!. Tem um exemplo no `arco2/server_v4/server_v4.go`

No arquivo `arco2/server_v4/server_gurl_v4.go` foi adicionado um sleep de 2s para analisarmos melhor como o servidor está se comportando, você pode adicionar a linha: `time.Sleep(2 * time.Second)`.<br>

Um resultado que talvez você não estivesse esperando, é que todas as requests são tratadas juntas, isso se dá por conta do listenAndServe, ele vai criar uma nova `goroutine` para cada uma das requests, as nossas funções de handler são executadas dentro de goroutines! Pense em goroutines como threads virtuais muito leves.

# Introduzindo rastreabilidade de requests

Um problema que vocês podem ter imaginado, é o volume de dados nos logs, isso cresce muito com a quantidade de usuários... Tornando quase impossível rastrear uma request, hoje tudo que temos no log é a mensagem e a rota que ela veio.<br>
Uma possível solução para isso é o que chamamos de requestId, é uma string, normalmente um UUID, que serve para identificar unicamente uma request, desse modo, conseguimos rastrear toda a request a partir de uma filtragem automática.<br>

Esse poder de rastreabilidade é muitíssimo importante no mundo real, imagine tentar descobrir o que causou um problema enfrentado por um cliente se seu servidor não tem nenhum tipo de log. Você estaria completamente de mãos atadas... Com os logs e a rastreabilidade com um requestId, nós temos o superpoder de descobrir exatamente o que aconteceu com uma request do passado.

Uma possível implementação seria algo assim:


```go
// server_v4

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
```

smallID é uma função que retorna uma string aleatória, a implementação dela também está no server_v4.

Com essas pequenas alterações, agora somos capazes de rastrear todo o lifecycle de uma request.
