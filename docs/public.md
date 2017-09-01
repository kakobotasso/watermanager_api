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
### Autenticação
#### Login [POST /v1/signin]
Verifica se o usuário e a senha existem no banco de dados e retorna o token de acesso

Parâmetros:

| Parâmetro        | Tipo           | Exemplo  |
| ---------------- |:--------------:| --------:|
| username         | `string`       | gopher   |
| password         | `string`       | 123456   |

Exemplo de resposta de sucesso:
```json
{
    "id": 123,
    "name": "Gopher",
    "token": "skjdfihs@#nsdj&jsdnfspai239uwe"
}
```

Exemplo de resposta de falha:
```json
[
    {
        "key": "login_incorrect",
        "message": "Username or password incorrect"
    }
]
```

### Usuários
#### Criação de usuário [POST /v1/user]
Cria um novo usuário

Parâmetros:

| Parâmetro        | Tipo           | Exemplo               |
| ---------------- |:--------------:| ---------------------:|
| name             | `string`       | Gopher                |
| email            | `string`       | gopher@golang.com     |
| cpf              | `string`       | 123.456.789-00        |
| username         | `string`       | gopher                |
| password         | `string`       | 123456                |

Exemplo de resposta:
```json
{
    "id": 123,
    "name": "Gopher",
    "token": "skjdfihs@#nsdj&jsdnfspai239uwe"
}
```

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