package util

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// GracefullyShutdown 优雅关闭 http server
func GracefullyShutdown(server *http.Server, gracePeriod time.Duration) {
	done := make(chan os.Signal, 1)
	/**
	os.Interrupt           -> ctrl+c 的信号
	syscall.SIGINT|SIGTERM -> kill 进程时传递给进程的信号
	*/
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-done
	LogrusObject.Println("closing http server gracefully...")

	ctx, cancel := context.WithTimeout(context.Background(), gracePeriod)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		LogrusObject.Fatalf("closing http server gracefully failed, err: %s", err.Error())
	}
}
