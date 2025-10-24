package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	url := obterUrl() // obtém a url do parâmetro do terminal

	resposta, err := http.Get(url) // A boa notícia é que GO já tem toda a parte complicada pronta, então podemos só usar!!
	if err != nil {
		fmt.Println("nao foi possível se comunicar com o servidor: ", err)
		os.Exit(1)
	}

	// Só fazemos ler a reposta e imprimir na tela
	bytesDaResposta := lerResposta(resposta)
	fmt.Println(bytesDaResposta)
}

func obterUrl() string {
	// os.Args tem todos os parâmetros que passamos mais o nome do arquivo na primeira posição
	if len(os.Args) < 2 {
		fmt.Println("Usagem inválida, passe uma url como parâmetro.")
		os.Exit(1)
	}
	return os.Args[1] // os.Args[0] é o nome do arquivo, então estamos interessados no segundo elemento
}

func lerResposta(resp *http.Response) string {
	body, err := io.ReadAll(resp.Body) // Aqui lemos a resposta do servidor para um array de bytes
	resp.Body.Close()                  // Precisamos sempre fechar as "streams" em GO (isso seria discussão para um outro workshop)

	if err != nil {
		fmt.Println("Não foi possível ler a resposta: ", err)
		os.Exit(1)
	}

	return string(body) // o que é uma string se não um array de bytes?
}
