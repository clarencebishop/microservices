package main
 
import (
    "github.com/clarencebishop/microservices/order/config"
    "github.com/clarencebishop/microservices/order/internal/adapters/db"
    "github.com/clarencebishop/microservices/order/internal/adapters/grpc"
    "github.com/clarencebishop/microservices/order/internal/application/core/api"
    "log"
)
 
func main() {
    dbAdapter, err := db.NewAdapter(config.GetDataSourceURL())
    if err != nil {
        log.Fatalf("Failed to connect to database. Error: %v", err)
    }
 
    application := api.NewApplication(dbAdapter)
    grpcAdapter := grpc.NewAdapter(application, config.GetApplicationPort())
    grpcAdapter.Run()
}
