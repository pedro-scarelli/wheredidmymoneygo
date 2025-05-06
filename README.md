# WhereDidMyMoneyGO

## Uma API RESTful simples em Go para rastreamento de despesas pessoais.
⚙️ Status: Em desenvolvimento !! – gerenciamento de usuários e autenticação já implementados.

### Recursos

    👤 Registro e gerenciamento de usuários

    🔐 Login baseado em JWT e rotas protegidas

    ➡️ Transferência entre contas

    📊 Planejado: CRUD de lançamentos de despesas, categorização, relatórios mensais

## Configuração e Instalação

    Clone o repositório

git clone https://github.com/pedro-scarelli/wheredidmymoneygo.git
cd wheredidmymoneygo

## Instale dependências

    go mod download

    Configure o ambiente
    Crie um arquivo .env na raiz do projeto (veja Variáveis de Ambiente abaixo).

## Executando o Servidor

### com hot reload (se usar Air / Fresh / Reflex)
air

### ou simplesmente:
go run cmd/server/main.go

Por padrão, a API escuta na porta 3000.
Variáveis de Ambiente

## Crie um .env com pelo menos:

PORT=3000
JWT_SECRET=sua_chave_secreta_jwt
DB_DSN="usuario=... senha=... host=... dbname=... port=..."

Endpoints da API

Todos os endpoints assumem http://localhost:3000 a menos que a variável PORT seja alterada.
Conta
Método	Rota	Auth	Descrição
GET	/account	Sem auth	Lista todas as contas
GET	/account/{id}	Bearer JWT	Detalha conta pelo ID
POST	/account	Sem auth	Cria uma nova conta de usuário
DELETE	/account/{id}	Sem auth	Deleta conta (sem proteção ainda)
Exemplo: Criar Conta

POST /account HTTP/1.1
Content-Type: application/json

{
  "firstName": "Pedro",
  "lastName": "Scarelli",
  "cpf": "10823027910",
  "email": "pvscarelli@gmail.com",
  "password": "Testee#1"
}

## Autenticação
Método	Rota	    Auth	        Descrição
POST	  /login	  Sem auth	    Autentica e retorna JWT
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

## Transferência
Método	Rota	Auth	Descrição
POST	/transfer	Sem auth	Transfere valor para outra conta
Exemplo: Transferência

POST /transfer HTTP/1.1
Content-Type: application/json

{
  "toAccount": 4187503,
  "amount":   131234
}
