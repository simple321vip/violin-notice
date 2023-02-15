package common

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run(eg *gin.Engine, srvName string, addr string, stop func()) {

	srv := &http.Server{Addr: addr, Handler: eg}

	go func() {
		log.Printf("%s running in %s \n", srvName, addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("%s", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)

	defer cancel()

	if stop != nil {
		stop()
	}

	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("%s ShutDown, cause by : %v", srvName, err)
	}

	select {
	case <-ctx.Done():
		log.Println("time out")
	}

	log.Println("Shutdown is successful!")

}
