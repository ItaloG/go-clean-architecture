FROM golang:1.21-bullseye

WORKDIR /usr/go

COPY . .

EXPOSE 8080
EXPOSE 8000
EXPOSE 50051

CMD [ "go", "run", "cmd/ordersystem/main.go", "cmd/ordersystem/wire_gen.go" ]