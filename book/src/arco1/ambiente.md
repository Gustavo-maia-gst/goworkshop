# Ambiente

## O que precisamos?

Precisamos basicamente do GO instalado na máquina e de um cliente HTTP para fazermos requisições. Vamos lá?

### Instalando o GO

A instalação do GO é bem simples, você pode instalar diretamente pelo binário seguindo a [documentação oficial](https://go.dev/doc/install).
Ou pelo gerenciador de pacotes da sua distribuição, por exemplo:

```bash
sudo apt install golang-go # Debian based
```
```bash
sudo pacman -S go # Arch based
```

Para testar se a instalação foi bem sucedida, rode o comando:

```bash
go version
```

Caso seja imprimido algum texto com a versão do GO, está tudo certo!

### Clonando o repositório

Vamos clonar o repositório do workshop para termos acesso ao código fonte dos exemplos.

```bash
git clone https://github.com/Gustavo-maia-gst/goworkshop.git
```
