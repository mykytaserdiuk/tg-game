package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/mykytaserdiuk/souptgbot/internal/db/mysql"
	"github.com/mykytaserdiuk/souptgbot/internal/soap/handler"
	repo "github.com/mykytaserdiuk/souptgbot/internal/soap/repository/mysql"
	"github.com/mykytaserdiuk/souptgbot/internal/soap/service"
	"github.com/mykytaserdiuk/souptgbot/pkg/web"
	"github.com/rs/cors"
	"github.com/spf13/viper"
)

func main() {

	err := godotenv.Load("pr.env")
	if err != nil {
		log.Println(err)
	}
	viper.AutomaticEnv()
	log.Println(os.Environ())

	cfg := mysql.Config{}
	err = cfg.LoadFromEnv()
	if err != nil {
		log.Fatal(err)
	}
	err = cfg.Validate()
	if err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()
	r.Use(web.FixContentType)

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT"},
	})

	dbPool, err := mysql.NewPool(&cfg)
	if err != nil {
		log.Fatal(err)
	}

	walletRepo := repo.NewWalletRepo()
	userRepo := repo.NewUserRepo()

	walletService := service.NewWalletService(dbPool, walletRepo, userRepo)
	userService := service.NewUserService(dbPool, walletRepo, userRepo)

	hand := c.Handler(r)
	_ = handler.New(r, walletService, userService)

	log.Println("Server starting....")
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Println("Server start")
	log.Fatal(http.ListenAndServe("0.0.0.0:"+port, hand))
}
