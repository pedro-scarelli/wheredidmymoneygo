# WhereDidMyMoneyGO

## Uma API RESTful simples em Go para rastreamento de despesas pessoais.

⚙️ Status: Em desenvolvimento !! – gerenciamento de usuários e autenticação já implementados.

### Recursos

    👤 Registro e gerenciamento de usuários

    🔐 Login baseado em JWT e rotas protegidas

    📊 Planejado: CRUD de lançamentos de despesas, categorização, relatórios mensais

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

POST /account HTTP/1.1
Content-Type: application/json

{
"firstName": "Pedro",
"lastName": "Scarelli",
"cpf": "{CPF_DO_USUARIO}",
"email": "pvscarelli@gmail.com",
"password": "Testee#1"
}

## Autenticação

Método Rota Auth Descrição
POST /login Sem auth Autentica e retorna JWT
Exemplo: Login

POST /login HTTP/1.1
Content-Type: application/json

{
"email": "pvscarelli@gmail.com",
"password": "Testee#1"
}

Retorno de sucesso:

{
"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9…"
}
