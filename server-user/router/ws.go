package router

import (
	"github.com/gin-gonic/gin"
	"github.com/mazezen/ginframe/server-user/handler"
)

// WsRouter websocket routers
func WsRouter(r *gin.Engine) {
	r.GET("/ws", handler.WsPing)
}
