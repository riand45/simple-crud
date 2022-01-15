package main

import (
	"log"
	"majoo/handler"
	"majoo/helper"
	"majoo/repository/product"
	"majoo/repository/outlet"
	"majoo/repository/user"
	"majoo/services/auth"
	"net/http"
	"strings"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/majoo?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	outletRepository := outlet.NewRepository(db)
	productRepository := product.NewRepository(db)

	userService := user.NewService(userRepository)
	outletService := outlet.NewService(outletRepository)
	productService := product.NewService(productRepository, outletRepository)
	authService := auth.NewService()

	userHandler := handler.NewUserHandler(userService, authService)
	outletHandler := handler.NewOutletHandler(outletService, authService)
	productHandler := handler.NewProductHandler(productService, authService)

	router := gin.Default()
	api := router.Group("/api/v1")

	api.POST("/register", userHandler.RegisterUser)
	api.POST("/login", userHandler.LoginUser)
	api.POST("/email_checkers", userHandler.CheckEmailAvailability)
	api.POST("/upload_avatars", authMiddleware(authService, userService), userHandler.UploadAvatar)

	api.GET("/outlets", authMiddleware(authService, userService), outletHandler.GetOutlets)
	api.GET("/outlet/:id", authMiddleware(authService, userService), outletHandler.GetOutlet)
	api.POST("/outlet", authMiddleware(authService, userService), outletHandler.CreateOutlet)
	api.PUT("/outlet/:id", authMiddleware(authService, userService), outletHandler.UpdateOutlet)
	api.DELETE("/outlet/:id", authMiddleware(authService, userService), outletHandler.DeleteOutlet)

	api.GET("/products/:id", authMiddleware(authService, userService), productHandler.GetProducts)
	api.POST("/product/:id", authMiddleware(authService, userService), productHandler.CreateProduct)
	api.POST("/upload_products", authMiddleware(authService, userService), productHandler.UploadImage)

	router.Use(cors.Default())
	router.Run()
}

func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		defaultToken := ""
		dataToken := strings.Split(authHeader, " ")
		if len(dataToken) == 2 {
			defaultToken = dataToken[1]
		}

		token, err := authService.ValidateToken(defaultToken)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, success := token.Claims.(jwt.MapClaims)

		if !success || !token.Valid {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userID := int(claim["user_id"].(float64))

		userSuccess, err := userService.GetUserById(userID)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Set("currentUser", userSuccess)
	}
}
