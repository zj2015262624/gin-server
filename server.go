package ginserver

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/go-oauth2/oauth2/v4"
	"github.com/go-oauth2/oauth2/v4/server"
)

var (
	gServer *server.Server
	once    sync.Once
)

// InitServer Initialize the service
func InitDefaultServer(manager oauth2.Manager) *server.Server {
	once.Do(func() {
		gServer = server.NewDefaultServer(manager)
	})
	return gServer
}
func InitServer(manager oauth2.Manager) *server.Server {
	once.Do(func() {
		conf := server.NewConfig()
		conf.ForcePKCE = true
		gServer = server.NewServer(conf, manager)
	})
	return gServer
}

// HandleAuthorizeRequest the authorization request handling
func HandleAuthorizeRequest(c *gin.Context) {
	err := gServer.HandleAuthorizeRequest(c.Writer, c.Request)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.Abort()
}

// HandleTokenRequest token request handling
func HandleTokenRequest(c *gin.Context) {
	err := gServer.HandleTokenRequest(c.Writer, c.Request)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.Abort()
}
