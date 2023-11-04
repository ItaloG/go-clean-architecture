# API para criar e listar pedidos
### Projeto criado usando a estrutura da Clean Architecture para se comunicar atraves de HTTP REST, gRpc e GraphQl 

O projeto utiliza de migrations criadas usando [golang-migrate](https://github.com/golang-migrate/migrate) e [docker](https://www.docker.com/) para criar containers do mysql e rabbitmq

antes de iniciar o projeto execute os seguintes comandos
- passo 1
```sh
  docker compose up -d
```
- passo 2
```sh
  make migrate ou migrate -path=sql/migrations -database "mysql://root:root@tcp(localhost:3306)/orders" -verbose up
```

Para rodar o projeto execute o comando a baixo dentro do diretório cmd/ordersystem

```sh
  go run main.go wire_gen.go
```

Para testar via HTTP use 

``Criar pedido``
```sh
    curl -X POST http://localhost:8000/order -H "Content-Type: application/json" -d "{"id":"a","price": 100.5,"tax": 0.5}"
```

``Listar pedidos``
```sh
    curl http://localhost:8000/order
```

Para testar via GraphQl playground

``Criar pedido``
>mutation createOrder {
  createOrder (input: {id: "graph_id", Price: 100.0, Tax: 0.2}) {
    id
    Price
    Tax
    FinalPrice
  }
>}

``Listar pedidos``
>query ListOrders {
  orders {
  id
  Price
  Tax
  FinalPrice
 }
>}

Para testar via gRpc utilizando o **evans** selecione o package *pb*, service *OrderService* e chame o *CreateOrder* para criar um pedido, ou *ListOrders* para listar pedidos