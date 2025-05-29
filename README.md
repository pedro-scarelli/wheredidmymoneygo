# WhereDidMyMoneyGO

## Uma API RESTful simples em Go para rastreamento de despesas pessoais.

⚙️ Status: Em desenvolvimento !! – gerenciamento de usuários e autenticação já implementados.

### Recursos

    👤 Registro e gerenciamento de usuários

    🔐 Login baseado em JWT e rotas protegidas

    📊 Planejado: Exportação de movimentos com saldo pra pranilha de Excel

## 📦 Pré-requisitos

- Docker
- Docker Compose
- Git

## Configuração e Instalação

Clone o repositório

```bash
git clone https://github.com/pedro-scarelli/wheredidmymoneygo.git
```

Navegue até a raiz do projeto

```bash
cd wheredidmymoneygo
```

## Executando o Servidor

```bash
docker compose up -d
```

Endpoints da API

URL base: http://localhost:3000

### Registrar conta

POST /account HTTP/1.1
Content-Type: application/json

{
"firstName": "{PRIMEIRO_NOME}",
"lastName": "{ULTIMO_NOME}",
"cpf": "{CPF}",
"email": "{EMAIL}",
"password": "{SENHA}"
}

### Buscar conta por ID

GET /account/{ID} HTTP/1.1
Authorization: Bearer {TOKEN_DO_USUARIO}

### Atualizar conta

PATCH /account HTTP/1.1
Content-Type: application/json
Authorization: Bearer {TOKEN_DO_USUARIO}

{
"id": "{ID}",
"firstName": "{PRIMEIRO_NOME}",
"lastName": "{ULTIMO_NOME}",
"password": "{SENHA}"
}

### Deletar conta por ID

DELETE /account/{ID} HTTP/1.1
Authorization: Bearer {TOKEN_DO_USUARIO}

## Autenticação

Método Rota Auth Descrição
POST /login HTTP/1.1
Content-Type: application/json

{
"email": "{EMAIL}",
"password": "{SENHA}"
}

Retorna token de autenticação

## Movimentação

POST /account/movement'
Authorization: Bearer {TOKEN_DO_USUARIO}

{
"value": {VALOR_COM_2_CASAS_APOS_A_VIRGULA},
"type": "{CREDITO_OU_DEBITO}"
"recurrence": {NUMERO_DE_RECORRENCIA - 1 A 12},
"dueDate": "{DATA_NO_FORMATO 2025-05-27T00:00:00Z}",
"description": "{DESCRICAO}"
}
