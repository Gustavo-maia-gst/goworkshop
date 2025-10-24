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
