package route

import (
	"assignment2go/config"
	"assignment2go/controller"
	"assignment2go/repository"
	"assignment2go/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Route(r *gin.Engine) {
	var (
		db              *gorm.DB                   = config.SetupDatabaseConnection()
		orderRepository repository.OrderRepository = repository.NewOrderRepository(db)
		orderService    service.OrderService       = service.NewOrderService(orderRepository)
		orderController controller.OrderController = controller.NewOrderController(orderService)
	)
	Route := r.Group("orders")
	{
		Route.GET("/", orderController.All)
		Route.POST("/", orderController.Insert)
		Route.PUT("/:id", orderController.Update)
		Route.DELETE("/:id", orderController.Delete)
	}
}
