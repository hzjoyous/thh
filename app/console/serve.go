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
	"runtime"
	"thh/arms/config"
	"thh/arms/logger"
	"thh/conf"
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
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Println("Thousand-hand:start")
	fmt.Printf("Thousand-hand:useMem %d KB\n", m.Alloc/1024/8)

	go RunJob()

	// 初始化应用程序
	if config.GetBool("app.debug", true) {
		go func() {
			// go tool pprof http://localhost:6060/debug/pprof/profile
			//http://127.0.0.1:7070/debug/pprof/
			err := http.ListenAndServe("0.0.0.0:7070", nil)
			if err != nil {
				fmt.Println(err)
			}
		}()
	}

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
