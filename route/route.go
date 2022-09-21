package route

import (
	"assignment2go/config"
	"assignment2go/controller"
	"assignment2go/middleware"
	"assignment2go/repository"
	"assignment2go/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Route(r *gin.Engine) {
	var (
		db              *gorm.DB                   = config.SetupDatabaseConnection()
		userRepository  repository.UserRepository  = repository.NewUserRepository(db)
		orderRepository repository.OrderRepository = repository.NewOrderRepository(db)
		jwtService      service.JWTService         = service.NewJWTService()
		authService     service.AuthService        = service.NewAuthService(userRepository)
		orderService    service.OrderService       = service.NewOrderService(orderRepository)
		userService     service.UserService        = service.NewUserService(userRepository)
		userController  controller.UserController  = controller.NewUserController(userService, jwtService)
		orderController controller.OrderController = controller.NewOrderController(orderService, jwtService)
		authController  controller.AuthController  = controller.NewAuthController(authService, jwtService)
	)
	Route := r.Group("orders", middleware.AuthorizeJWT(jwtService))
	{
		Route.GET("/", orderController.All)
		Route.POST("/", orderController.Insert)
		Route.PUT("/:id", orderController.Update)
		Route.DELETE("/:id", orderController.Delete)
	}

	AuthRoute := r.Group("auth")
	{
		AuthRoute.POST("/register", authController.Register)
		AuthRoute.POST("/login", authController.Login)
	}

	UserRoute := r.Group("user", middleware.AuthorizeJWT(jwtService))
	{
		UserRoute.GET("/", userController.FindUser)
		UserRoute.PUT("/", userController.Update)
	}

}
