#DATA_SOURCE_URL=postgres:postgres@tcp(127.0.0.1:5432)/order APPLICATION_PORT=3000 ENV=development go run cmd/main.go
DATA_SOURCE_URL=postgres:postgres@127.0.0.1:5432/order APPLICATION_PORT=3000 ENV=development go run cmd/main.go
