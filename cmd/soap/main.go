package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/mykytaserdiuk/souptgbot/internal/db/mysql"
	"github.com/mykytaserdiuk/souptgbot/internal/soap/handler"
	"github.com/rs/cors"
	"github.com/spf13/viper"
)

func fixContentType(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		log.Println(r.RequestURI)
		w.Header().Add("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")

		next.ServeHTTP(w, r)
	})
}

func main() {

	log.Println(os.Environ())
	viper.AutomaticEnv()

	cfg := mysql.Config{}
	err := cfg.LoadFromEnv()
	if err != nil {
		log.Fatal(err)
	}
	err = cfg.Validate()
	if err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()
	r.Use(fixContentType)
	_ = handler.New(r)

	log.Println("Server starting....")
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT"},
	})

	hand := c.Handler(r)
	_, err = mysql.NewPool(&cfg)
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(http.ListenAndServe("0.0.0.0:"+port, hand))
	log.Println("Server start")
}
