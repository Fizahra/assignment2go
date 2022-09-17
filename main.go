package main

import (
	"assignment2go/config"
	"assignment2go/controller"
	"assignment2go/repository"
	"assignment2go/service"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	var (
		db              *gorm.DB                   = config.SetupDatabaseConnection()
		orderRepository repository.OrderRepository = repository.NewOrderRepository(db)
		// itemRepository repository.ItemRepository = repository.NewItemRepository(db)
		orderService service.OrderService = service.NewOrderService(orderRepository)
		// authController controller.AuthController = controller.NewAuthController()
		orderController controller.OrderController = controller.NewOrderController(orderService)
	)
	serverAddress := os.Getenv("SERVICE_PORT")
	r := gin.Default()
	Route := r.Group("orders")
	{
		Route.GET("/", orderController.All)
		Route.POST("/", orderController.Insert)
		Route.PUT("/:id", orderController.Update)
		Route.DELETE("/:id", orderController.Delete)
	}
	r.Run(serverAddress)
}
