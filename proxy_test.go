package goutils

import (
	"github.com/gin-gonic/gin"
	"os"
	"os/signal"
	"syscall"
	"testing"
	"time"
)

func TestNewProxyRouter(t *testing.T) {
	configs := []*ProxyConfig{
		NewProxyConfig(
			[]string{"http://127.0.0.1:8091", "http://127.0.0.1:8092"},
			"/api",
			WithHealthInterval(10),
			WithHealthCheck(true),
		),
		NewProxyConfig(
			[]string{"http://127.0.0.1:8081"},
			"/api",
			WithHealthCheck(false),
			WithReWritePath("/t"),
		),
	}

	g := gin.Default()
	root := g.Group("")
	InitProxyRouter(configs, root)
	go g.Run(":8080")
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	sig := <-signalChan
	if sig == syscall.SIGINT || sig == syscall.SIGTERM {
		time.Sleep(time.Second)
		os.Exit(0)
	}
}
