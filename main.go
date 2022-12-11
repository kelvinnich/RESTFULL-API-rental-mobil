package main

import (
	"restfull-api-rental-mobil/config"
	"restfull-api-rental-mobil/controller"
	entitycontroller "restfull-api-rental-mobil/entity/entity_controller"
	entityrepository "restfull-api-rental-mobil/entity/entity_repository"
	entityservice "restfull-api-rental-mobil/entity/entity_service"
	
	"restfull-api-rental-mobil/repository"
	"restfull-api-rental-mobil/service"
	"restfull-api-rental-mobil/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = config.ConnectDB()

	//repository
	userRepository entityrepository.UserRepository = repository.ConnectCustomerRepository(db)
	menuRepository entityrepository.MenuRepository = repository.ConnectMenuRepository(db)
	transactionRepository entityrepository.TransactionRepository = repository.ConnecetTransactionRepository(db)

	//service
	jwtService entityservice.JwtService = service.NewJwtService()
	authService entityservice.AuthService = service.NewConnectAuthService(userRepository)
	customerService entityservice.UserService = service.NewConnectUserService(userRepository)
	menuService entityservice.MenuService = service.NewConnectMenuService(menuRepository)
	transactionService entityservice.TransactionService = service.NewConnectTransactionsService(transactionRepository, menuRepository)

	//controller
	authController entitycontroller.AuthController = controller.NewVerifyController(authService, jwtService)
	userController entitycontroller.UserController = controller.NewConnectUserController(customerService, jwtService)
	menuController entitycontroller.MenuController = controller.NewMenuController(menuService, jwtService)
	transactionController entitycontroller.TransactionController = controller.NewTransactionController(transactionService, jwtService)
)

func main(){

	defer config.CloseDB(db)
	r := gin.Default()

	authRoutes := r.Group("api/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}

	customerRoutes := r.Group("api/user", middleware.AuthorizeJWT(jwtService))
	{
		customerRoutes.GET("/profile", userController.Profiles)
		customerRoutes.PUT("/profile", userController.Updates)
	}

	menuRoutes := r.Group("api/menu")
	{
		menuRoutes.GET("/", menuController.All)
		menuRoutes.POST("/", menuController.Insert)
		menuRoutes.GET("/:id", menuController.FindMenuByID)
		menuRoutes.PUT("/:id", menuController.Update)
		menuRoutes.DELETE("/:id", menuController.Delete)

	}

	transactionRoutes := r.Group("api/transaction")
	{
		transactionRoutes.POST("/", transactionController.InsertTx)
		transactionRoutes.GET("/", transactionController.AllTx)
		transactionRoutes.PUT("/:id", transactionController.UpdateTx)
		transactionRoutes.DELETE("/:id", transactionController.DeleteTx)
	}

	r.Run(":8000")
}