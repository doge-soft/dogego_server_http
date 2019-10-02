package dogego_server_http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

// HTTP/1.1 服务器
func HTTPServerProtocol(router *gin.Engine) error {
	err := http.ListenAndServe(os.Getenv("ADDR_HTTP"), router)

	if err != nil {
		return err
	}

	return nil
}
