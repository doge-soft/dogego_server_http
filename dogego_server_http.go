package dogego_server_http

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

// HTTP/1.1 服务器
func HTTPServerProtocol(router *gin.Engine) error {
	log.Printf("HTTP/1.1 Server started on %s", os.Getenv("ADDR_HTTP"))
	err := http.ListenAndServe(os.Getenv("ADDR_HTTP"), router)

	if err != nil {
		return err
	}

	return nil
}
