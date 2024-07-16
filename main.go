package main

import (
	"fmt"
	"log"
	"net/http"
	"pln/auth"
	"pln/handler"
	"pln/helper"
	"pln/level"
	"pln/pelanggan"
	"pln/pembayaran"
	"pln/penggunaan"
	"pln/tagihan"
	"pln/tarif"
	"pln/user"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// koneksi ke database PLN
	dsn := "root:@tcp(127.0.0.1:3306)/pln?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	// load env
	err = godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }

	// repository
	userRepository := user.NewRepostiory(db)
	penggunaanRepository := penggunaan.NewRepository(db)
	tarifRepository := tarif.NewRepository(db)
	tagihanRepository := tagihan.NewRepository(db)
	pelangganRepository := pelanggan.NewRepository(db)
	pembayaranRepository := pembayaran.NewRepository(db)
	levelRepository := level.NewRepository(db)
	// service
	userService := user.NewService(userRepository, pelangganRepository)
	penggunaanService := penggunaan.NewService(penggunaanRepository)
	tarifService := tarif.NewService(tarifRepository)
	tagihanService := tagihan.NewService(tagihanRepository)
	authService := auth.NewService()
	pelangganService := pelanggan.NewService(pelangganRepository)
	pembayaranService := pembayaran.NewService(pembayaranRepository)
	levelService := level.NewService(levelRepository)
	// handler
	userHandler := handler.NewUserHandler(userService, authService)
	penggunaanHandler := handler.NewPenggunaanHandler(penggunaanService)
	tarifHandler := handler.NewTarifHandler(tarifService)
	tagihanHandler := handler.NewTagihanHandler(tagihanService)
	pelangganHandler := handler.NewPelangganHandler(pelangganService)
	pembayaranHandler := handler.NewPembayaranHandler(pembayaranService)
	levelHandler := handler.NewLevelHandler(levelService)

	// router
	router := gin.Default()
	api := router.Group("/api/v1")

	// router private
	private := router.Group("/private/api/v1").Use(authMiddleware(authService, userService, pelangganService))

	// api
		// auth
	api.POST("/user", userHandler.RegisterUser)
	api.POST("/login", userHandler.LoginUser)
		// penggunaan
	private.POST("/create-penggunaan", penggunaanHandler.CreatePenggunaan)
	private.GET("/penggunaan", penggunaanHandler.GetPenggunaan)
	private.PATCH("/edit-penggunaan/:id_penggunaan", penggunaanHandler.UpdatePenggunaan)
	private.DELETE("/delete-penggunaan/:id_penggunaan", penggunaanHandler.DeletePenggunaan)
		// tarif
	private.POST("/create-tarif", tarifHandler.CreateTarif)
	private.GET("/tarif", tarifHandler.GetTarif)
	private.PATCH("/edit-tarif/:id_tarif", tarifHandler.UpdateTarif)
	private.DELETE("delete-tarif/:id_tarif", tarifHandler.DeleteTarif)
		// tagihan
	private.POST("/create-tagihan", tagihanHandler.CreateTagihan)
	private.GET("/tagihan", tagihanHandler.GetTagihan)
	private.PATCH("/edit-tagihan/:id_tagihan", tagihanHandler.UpdateTagihan)
	private.DELETE("/delete-tagihan/:id_tagihan", tagihanHandler.DeleteTagihan)
		// pelanggan
	private.POST("/create-pelanggan", pelangganHandler.CreatePelanggan)
	private.GET("/pelanggan", pelangganHandler.GetPelanggan)
		// pembayaran
	private.POST("/create-pembayaran", pembayaranHandler.CreatePembayaran)
	private.GET("/pembayaran", pembayaranHandler.GetPembayaran)
		// level
	private.POST("/create-level", levelHandler.CreateLevel)
	private.GET("/level", levelHandler.GetLevel)

	router.Run()
}


func authMiddleware(authService auth.Service, userService user.Service, pelangganService pelanggan.Service) gin.HandlerFunc {
	return func (c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if !strings.Contains(authHeader, "Bearer") {
			fmt.Println("SSS")
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}
		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			fmt.Println("====")
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		claim, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid{
			fmt.Println("???")
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		fmt.Println(claim,"claimnya")
		
		userID := fmt.Sprintf("%v", claim["id_user"])
		user, err := userService.GetUserByID(userID)
		fmt.Println(user,"user")

		if err != nil && strings.Contains(fmt.Sprint(err), "No user found on with that ID") {			
			pelanggans, err := pelangganService.GetFindID(userID)
			if err != nil {
				response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
				c.AbortWithStatusJSON(http.StatusUnauthorized, response)
				return
			}
			user.IDUser = pelanggans.IDPelanggan
			
		}
		c.Set("currentUser", user)
	}
}