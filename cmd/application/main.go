package main

import (
	"fmt"
	_ "github.com/3tonColl/music_task/docs"
	"github.com/3tonColl/music_task/internal/config"
	"github.com/3tonColl/music_task/internal/controller"
	"github.com/3tonColl/music_task/internal/dbController"
	"github.com/3tonColl/music_task/internal/handler"
	"github.com/3tonColl/music_task/internal/router"
	"github.com/3tonColl/music_task/logger"
	"github.com/joho/godotenv"
	"strconv"
)

func main() {
	var lgr = logger.NewLogger()
	err := godotenv.Load("../.env")
	if err != nil {
		lgr.ErrLogger.Fatalf("Error loading .env file")
	} else {
		lgr.InfoLogger.Printf(".env file was found and loaded")
	}

	defer func() {
		if rec := recover(); rec != nil {
			lgr.ErrLogger.Printf("Caught panic: %v", rec)
		}
	}()
	conf := config.NewConfig()
	connectionStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", conf.DB.DB_USER, conf.DB.DB_PASSWORD, conf.DB.DB_HOST, strconv.Itoa(conf.DB.DB_PORT), conf.DB.DB_NAME)

	songRepo, err := dbController.NewdbController(connectionStr, lgr)
	if err != nil {
		lgr.ErrLogger.Fatalf("Failed to create a dbController instance.")
	}
	songController := controller.NewSongController(songRepo, lgr)
	songHandler := handler.NewSongHandler(songController, lgr)

	if err != nil {
		panic(fmt.Errorf("Initialization has failed: %s\n", err))
	}
	lgr.InfoLogger.Println("Initialization components for router has successfully")
	app := router.NewRouter(songHandler)
	lgr.DebugLogger.Println("Launching the application.....")
	app.Listen(fmt.Sprintf(":%s", strconv.Itoa(conf.API.API_PORT)))
}
