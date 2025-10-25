package main

import (
	"io"
	"net/http"
	"time"
)

// Esse é o mesmo código do GURL que fizemos encapsulado numa função retornando os bytes ao invés de imprimir eles
// copiado e colado de arco1/gurl.go

func enviarRequisicao(url string) ([]byte, error) {
	resposta, err := http.Get(url) // A boa notícia é que GO já tem toda a parte complicada pronta, então podemos só usar!!
	if err != nil {
		return nil, err
	}

	time.Sleep(2 * time.Second) // pausa 2 segundos

	// Só fazemos ler a reposta e imprimir na tela
	bytesDaResposta, err := lerResposta(resposta)
	if err != nil {
		return nil, err
	}

	return bytesDaResposta, nil
}

func lerResposta(resp *http.Response) ([]byte, error) {
	body, err := io.ReadAll(resp.Body) // Aqui lemos a resposta do servidor para um array de bytes
	resp.Body.Close()                  // Precisamos sempre fechar as "streams" em GO (isso seria discussão para um outro workshop)

	if err != nil {
		return nil, err
	}

	return body, nil // o que é uma string se não um array de bytes?
}
