# WaterManager API

## Objetivo
A WaterManager API é uma aplicação que será responsável pela conversa entre o medidor de fluxo de água do Arduino e o aplicativo mobile WaterManager

## Documentação
A documentação dos endpoints da API pode ser encontrados [clicando aqui](https://github.com/kakobotasso/watermanager_api/blob/master/docs/public.md)

## Rodando o projeto na sua máquina
1. Baixar o Go [clicando aqui](https://golang.org/dl/) e seguir o wisard de instalação;
2. Abrir o arquivo `$HOME/.bash_profile` e adicionar: 
    ```
    export GOROOT=/usr/local/go
    export GOPATH=$HOME/go
    export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
    ```
4. Rodar os comandos:
    ```
    source $HOME/.bash_profile
    mkdir -p $HOME/go/{bin,pkg,src/github.com/kakobotasso}
    cd $HOME/go/src/github.com/kakobotasso
    ```
5. Clonar o projeto dentro dessa pasta;
6. Para subir a aplicação, rode o comando:
    ```
    go run main.go
    ```
7. Abra o [Postman](https://www.getpostman.com/) e de um GET em:
    ```
    http://localhost:3000/healthcheck
    ```