# Clean Architecture: Listagem de Orders

Listagem de orders com um único use case (`ListOrdersUseCase`) exposto via REST, gRPC e GraphQL.

## Como executar

```bash
docker compose up
```

O comando sobe o MySQL e o RabbitMQ, aplica as migrações e inicia a aplicação.

## Portas

| Serviço | Porta | Acesso |
|---|---|---|
| Web (REST) | 8000 | `POST /order` e `GET /order` |
| gRPC | 50051 | `OrderService`: `CreateOrder` e `ListOrders` |
| GraphQL | 8080 | Playground em `/` e endpoint em `/query` |

## Como testar

### REST e GraphQL

O arquivo `api.http` tem as requisições para criar e listar orders.

### GraphQL

Em http://localhost:8080:

```graphql
query {
  ListOrders {
    id
    Price
    Tax
    FinalPrice
  }
}
```

### gRPC

O servidor tem reflection habilitado. Com o Evans:

```bash
evans -r repl
```

```
package pb
service OrderService
call ListOrders
```

## Migrações

Os arquivos ficam em `sql/migrations` e são aplicados pelo serviço `migrate` do `docker-compose.yaml`.
A aplicação só inicia depois que o MySQL está saudável e as migrações terminam com sucesso.

## Testes

```bash
go test ./...
```

Os testes do repositório usam SQLite e precisam de CGO habilitado.
