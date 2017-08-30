# WaterManager Endpoints

# Aplicação
#### Health Check [GET /healthcheck]
Verifica o status da aplicação

Exemplo de resposta:
```json
{ "status": 200 }
```

#### Version [GET /version]
Verifica a versão do app

Exemplo de resposta:
```json
{ "version": "0.1.0" }
```
***

# /v1
### Condomínios
### Usuários
#### Informações do usuário [GET /v1/user/:id]
Recupera as informações de um usuário

Parâmetros:

| Parâmetro        | Tipo           | Exemplo  |
| ---------------- |:--------------:| --------:|
| id               | `int`          | 123      |

Exemplo de resposta:
```json
{
    "id": 123,
    "name": "Gopher",
    "email": "gopher@golang.com",
    "cpf": "123.456.789-00",
    "username": "gopher",
    "password": "123456",
    "token": "skjdfihs@#nsdj&jsdnfspai239uwe"
}
```