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

### Buscar todas as contas
GET /account?page={PAGINA}&temsPerPage={ITENS_POR_PAGINA} HTTP/1.1

### Buscar conta por ID
GET /account/{ID} HTTP/1.1

### Atualizar conta
PATCH /account HTTP/1.1
Content-Type: application/json

{
    "id": "{ID}",
    "firstName": "{PRIMEIRO_NOME}",
    "lastName": "{ULTIMO_NOME}",
    "password": "{SENHA}"
}

### Deletar conta por ID
DELETE /account/{ID} HTTP/1.1

## Autenticação

Método Rota Auth Descrição
POST /login HTTP/1.1
Content-Type: application/json

{
"email": "{EMAIL}",
"password": "{SENHA}"
}

Retorna token de autenticação
