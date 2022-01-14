package console

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"log"
	"net/http"
	"os"
	"os/signal"
	"thh/conf"
	"thh/helpers/config"
	"thh/helpers/logger"
	"thh/routes"
	"time"
)

// CmdServe represents the available web sub-command.
var CmdServe = &cobra.Command{
	Use:   "serve",
	Short: "Start web server",
	Run:   runWeb,
	Args:  cobra.NoArgs,
}

func runWeb(cmd *cobra.Command, args []string) {

	go RunJob()

	port := config.GetString("app.port")
	var engine *gin.Engine
	switch conf.IsProd() {
	case true:
		gin.DisableConsoleColor()
		gin.SetMode(gin.ReleaseMode)
		engine = gin.New()
		break
	default:
		engine = gin.Default()
		break
	}

	routes.RegisterRoutes(engine)

	srv := &http.Server{
		Addr:           ":" + port,
		Handler:        engine,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	fmt.Println("Thousand-hand:listen " + port)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	// 这一句是 go chan 等待接受值，只是把接到的值直接扔掉了，此处是主协程的阻塞处
	_ = <-quit

	logger.Std().Println("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Std().Println("Server Shutdown:", err)
	}
	logger.Std().Println("Server exiting")
}
