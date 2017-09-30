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
    "created_at": "0001-01-01T00:00:00Z",
    "updated_at": "0001-01-01T00:00:00Z",
    "name": "Gopher",
    "token": "skjdfihs@#nsdj&jsdnfspai239uwe",
    "serial": "123aosdna423"
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
| serial           | `string`       | 123aosdna423          |

Exemplo de resposta:
```json
{
    "created_at": "0001-01-01T00:00:00Z",
    "updated_at": "0001-01-01T00:00:00Z",
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
    "created_at": "2017-09-26T00:46:43.185414607-03:00",
    "updated_at": "2017-09-26T00:46:43.185414607-03:00",
    "name": "Gopher",
    "email": "gopher@golang.com",
    "cpf": "123.456.789-00",
    "username": "gopher",
    "password": "123456",
    "token": "skjdfihs@#nsdj&jsdnfspai239uwe",
    "serial": "123aosdna423"
}
```

### Consumo
#### Criação de consumo [POST /v1/consumption]
Cria um novo consumo

Parâmetros:

| Parâmetro        | Tipo           | Exemplo               |
| ---------------- |:--------------:| ---------------------:|
| liter            | `int`          | 12                    |
| month            | `int`          | 09                    |
| year             | `int`          | 2017                  |
| serial           | `string`       | 123aosdna423          |

Exemplo de resposta:
```json
{
    "id": 1,
    "liter": "2",
    "created_at": "2017-09-29T00:50:53.72423515-03:00",
    "updated_at": "2017-09-29T00:50:53.72423515-03:00",
    "month": "09",
    "year": "2017",
    "serial": "123aosdna423"   
}
```

Exemplo de resposta de falha:
```json
[
    {
        "key": "consumption_create_failure",
        "message": "Unable to create a new consumption"
    }
]
```

#### Lista de consumo [GET /v1/consumption/:serial/:consumption_type]
Recupera uma lista de com o consumo do usuário

Parâmetros:

| Parâmetro                | Tipo           | Exemplo         |
| ------------------------ |:--------------:| ---------------:|
| serial                   | `string`       | 123aosdna423    |
| consumption_type         | `string`       | liter / money   |

Exemplo de resposta de sucesso:
```json
[
    {
        "id": 1,
        "created_at": "2017-09-29T00:50:53.72423515-03:00",
        "updated_at": "2017-09-29T00:50:53.72423515-03:00",
        "liter": "2",
        "month": "09",
        "year": "2017",
        "serial": "123aosdna423"
    },
    {
        "id": 2,
        "created_at": "2017-09-29T00:50:53.72423515-03:00",
        "updated_at": "2017-09-29T00:50:53.72423515-03:00",
        "liter": "2",
        "month": "08",
        "year": "2017",
        "serial": "123aosdna423"
    }
]
```

Exemplo de resposta de falha:
```json
[
    {
        "key": "invalid_parameter",
        "message": "Invalid type of consumption"
    }
]
```

#### Lista de consumo mensal [GET /v1/consumption/monthly/:serial/:consumption_type]
Recupera uma lista de com o consumo mensal do usuário

Parâmetros:

| Parâmetro                | Tipo           | Exemplo         |
| ------------------------ |:--------------:| ---------------:|
| serial                   | `string`       | 123aosdna423    |
| consumption_type         | `string`       | liter / money   |

Exemplo de resposta de sucesso:
```json
[
    {
        "id": 1,
        "created_at": "2017-09-29T00:50:53.72423515-03:00",
        "updated_at": "2017-09-29T00:50:53.72423515-03:00",
        "liter": "2",
        "month": "09",
        "year": "2017",
        "serial": "123aosdna423"
    },
    {
        "id": 2,
        "created_at": "2017-09-29T00:50:53.72423515-03:00",
        "updated_at": "2017-09-29T00:50:53.72423515-03:00",
        "liter": "2",
        "month": "08",
        "year": "2017",
        "serial": "123aosdna423"
    }
]
```

#### Consumo estimado [GET /v1/consumption/estimated/:serial/:consumption_type]
Recupera a estimativa de consumo do usuário

Parâmetros:

| Parâmetro                | Tipo           | Exemplo         |
| ------------------------ |:--------------:| ---------------:|
| serial                   | `string`       | 123aosdna423    |
| consumption_type         | `string`       | liter / money   |

Exemplo de resposta de sucesso:
```json
[
    {
        "id": 1,
        "created_at": "2017-09-29T00:50:53.72423515-03:00",
        "updated_at": "2017-09-29T00:50:53.72423515-03:00",
        "liter": "2",
        "month": "09",
        "year": "2017",
        "serial": "123aosdna423"
    }
]
```

Exemplo de resposta de falha:
```json
[
    {
        "key": "invalid_parameter",
        "message": "Invalid type of consumption"
    }
]
```

#### Consumo estimado e mensal [GET /v1/consumption/monthly/estimated/:serial/:consumption_type]
Recupera a estimativa de consumo do usuário e seu consumo mensal

Parâmetros:

| Parâmetro                | Tipo           | Exemplo         |
| ------------------------ |:--------------:| ---------------:|
| serial                   | `string`       | 123aosdna423    |
| consumption_type         | `string`       | liter / money   |

Exemplo de resposta de sucesso:
```json
[
    {
        "estimated": [
            {
                "id": 1,
                "created_at": "0001-01-01T00:00:00Z",
                "updated_at": "0001-01-01T00:00:00Z",
                "liter": "2",
                "month": "9",
                "year": "2017",
                "serial": "123aosdna423"
            }
        ],
        "consumption_list": [
            {
                "id": 1,
                "created_at": "0001-01-01T00:00:00Z",
                "updated_at": "0001-01-01T00:00:00Z",
                "liter": "2",
                "month": "9",
                "year": "2017",
                "serial": "123aosdna423"
            }
        ]
    }
]
```

Exemplo de resposta de falha:
```json
[
    {
        "key": "invalid_parameter",
        "message": "Invalid type of consumption"
    }
]
```