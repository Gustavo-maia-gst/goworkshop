# Sua primeira request

# Noções básicas de redes

### Interpretação dummy para comunicação HTTP

De forma muito, mas muito, simplificada, pode-se pensar em requisições HTTP como uma forma de um computador (client) executar um procedimento remotamente por meio de um formato pré-definido (request) em outro computador (host) e obter um resultado (response).<br>

No nosso caso, o procedimento será executar função em GO e a resposta seu retorno, mas poderia ser retornar o conteúdo de um arquivo ou realizar qualquer operação que você possa imaginar...

### DNS e portas: Nomes na internet

Toda vez que você acessa um site como `www.google.com` ou aquele que não deve ser nomeado (`sigaa.ufcg.edu.br`), por exemplo, o que acontece é que o seu computador, por meio do *DNS* (outra sigla que vamos tratar como caixa preta) realiza a tradução desse nome para um endereço IP (192.0.2.0 ou algo assim) que pode ser usado para encontrar um computador acessível em algum lugar do mundo.<br>

Então o seu computador envia uma *requisição* para o computador do endereço encontrado direcionada a uma porta, sim, precisa estar direcionado a uma porta.<br>

No nosso caso, o servidor estará rodando na nossa própria máquina, o nome será localhost. A porta você pode escolher seu número da sorte (desde que ele seja maior que 1024, essas costumam ter uma semântica pré-definida - por exemplo, porta 80 é a porta HTTP- e requerem níveis maiores de privilégio para serem usadas).<br>
Para acessarmos o serviço rodando na máquina local na porta 8000, usamos localhost:8000.

<br>
<iframe width="560" height="315" src="https://www.youtube.com/embed/1K0swnzs25g?si=pF0_V0tBO4UOSGQ4" title="YouTube video player" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share" referrerpolicy="strict-origin-when-cross-origin" allowfullscreen></iframe>
<br>
<br>

> Não vamos se aprofundar mais que isso em HTTP ou redes, mas caso você se interesse em conhecer mais, recomendo os capítulos sobre HTTP e TCP do livro do Kurose (Computer Networking: A Top-Down Approach)

# "Watch and learn before you do." - Se comunicando com outros servidores

Nesta pequena seção vamos realizar algumas chamadas para APIs externas para tentar visualizar o que acontece e ter uma noção do que vamos construir.<br>
Vamos escolher alguma API pública para fazer algumas chamadas, eu selecionei algumas legais, você pode escolher a que te chamar mais atenção, ou qualquer outra:

| API | Descrição |
|-----|-----------|
| https://api.chucknorris.io/jokes/random | Retorna uma piada aleatória sobre o Chuck Norris. |
| https://rickandmortyapi.com/api/character/1 | Retorna dados sobre o Rick do Rick and Morty. |
| https://pokeapi.co/api/v2/pokemon/pikachu | Retorna dados sobre o Pikachu. |

> Para testar no terminal, podemos usar o comando `curl`, ficaria: `curl <url>`, se você não gostar muito do que ver e tiver o `jq` instalado, pode rodar `curl <url> | jq` e ver um resultado um pouco mais amigável.

### gurl: Vamos fazer uma pequena cópia do curl, recebendo uma url e imprimindo o resultado da chamada na tela!

O código ficaria assim:

```go
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

	// Esse é o padrão de tratamento de erro em GO, as funções retornam um objeto de erro e você precisa checar ele!
	// Pode parecer um pouco verboso no começo mais deixa tudo mais explicíto e previsível, qualidades muito subestimadas em um programa.
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
```
<br>

Você pode brincar um pouco com diferentes URLs, experimente o que acontece se passar uma url que não existe, ou tente trocar a URI (parte dps do .com/), o que será que acontecerá se você trocar o /pikachu por /squirtle? Ou o /1 do rick and morty por /2?<br>

> Repare no início da url `https://`, esse é o indicador do protocolo que deve ser usado, é recomendado sempre prefixar com o protocolo. O s após http significa safe, indica que é uma variante com criptografia do http, nos nossos exemplos em localhost, vamos usar `http://`
