# API para criar e listar pedidos
### Projeto criado usando a estrutura da Clean Architecture para se comunicar atraves de HTTP REST, gRpc e GraphQl 

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