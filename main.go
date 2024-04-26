package main

import (
	"fmt"
	"gin-api/app/injector"
	"gin-api/app/router"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	port := fmt.Sprintf(":%s", os.Getenv("APP_PORT"))

	init, err := injector.Init()
	if err != nil {
		panic(err)
	}
	db, err := init.DB.DB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	readTimeout, err := strconv.Atoi(os.Getenv("APP_READ_TIMEOUT"))
	if err != nil {
		panic(err)
	}
	writeTimeout, err := strconv.Atoi(os.Getenv("APP_WRITE_TIMEOUT"))
	app := &http.Server{
		Addr: port,
		Handler: router.Init(init),
		ReadTimeout: time.Duration(readTimeout) * time.Second,
		WriteTimeout: time.Duration(writeTimeout) * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	
	if err := app.ListenAndServe(); err != nil {
		panic(err)
	}
}