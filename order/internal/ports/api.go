package ports
 
import "github.com/clarencebishop/microservices/order/internal/application/core/domain"
 
type APIPort interface {
    PlaceOrder(order domain.Order) (domain.Order, error)
}
