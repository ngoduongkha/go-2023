package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ngoduongkha/go-2023/assignment-06/config"
	"github.com/ngoduongkha/go-2023/assignment-06/handlers"
	"github.com/ngoduongkha/go-2023/assignment-06/router"
)

func init() {
	config.InitializeAppConfig()
	if !config.AppConfig.Debug {
		gin.SetMode(gin.ReleaseMode)
	}
}

func main() {
	if err := handlers.InitializeHandler(); err != nil {
		log.Fatalln(err)
	}
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", config.AppConfig.Port),
		Handler:        router.InitializeRouter(),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	if err := s.ListenAndServe(); err != nil {
		log.Fatalln(err)
	}
}
