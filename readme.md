# Product Comparison API (Go + Gin)

API simples para **comparação de produtos** a partir de uma base JSON carregada em memória.
Ela recebe uma lista de IDs e (opcionalmente) uma lista de campos a comparar (`fields`).
Campos podem ser **top-level** do `Product` (ex: `price`, `size`) ou **dinâmicos** dentro de `specs` (ex: `brand`, `model`, `sleeve`).

## Endpoints

- `GET /` → health check
- `GET /ping` → retorna `pong`
- `GET /v1/compare/products?ids=1,2,3&fields=price,size,brand`
    - `ids` (**obrigatório**): lista separada por vírgula
    - `fields` (opcional): lista separada por vírgula. Se vazio, usa defaults do model.
    - Resposta retorna:
        - `requestedIds`, `found`, `notFound`
        - `comparison.fields` e `comparison.items` (lista por produto)
        - `summary` (requested/found/notFound/duplicated)

## Setup / execução

### Pré-requisitos
- Go instalado (compatível com `go.mod`)

### Configuração
- O arquivo `.env` define onde está a base:
    - `FILE_PATH=/projects/challenge/data/products.json`

### Rodar a aplicação
- `go run cmd/main.go`
- curl -i -H "X-Trace-Id: trace-123" "http://127.0.0.1:8080/v1/compare/products?ids=1,2,3"


### Rodar testes
- `go test ./...`

## Decisões de arquitetura

- **Gin** como framework HTTP.
- **Carregamento em memória**: a base de dados é via arquivo JSON que é lido na inicialização e mantido em memória.

**Camadas**:

**handler**
- validar parâmetros da requisição
- aplicar timeout de requisição (3s)
- retornar códigos HTTP adequados

### service
- implementar a lógica de comparação
- construir o payload de resposta
- aplicar regras de negócio

### repository
- buscar produtos pelos IDs
- identificar produtos `notFound`
- identificar IDs duplicados

### models
- definição da struct `Product`
- campos padrão de comparação (`GetProductDefaultFields`)
- **Trace/Logs**: header `X-Trace-Id` é propagado e usado em logs.
- **Interfaces**: Utilizada para desacoplar as camadas e deixar as dependências explícitas 

## TODOs
- Docker / docker-compose para facilitar execução
- cache em Redis para grandes volumes de produtos
- paginação de resultados
- métricas e observabilidade
- validação mais robusta de campos solicitados
