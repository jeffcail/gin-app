package boot

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mazezen/ginframe/server-user/global"
	"github.com/mazezen/ginframe/server-user/router"
)

// HttpServe init
func HttpServe() error {
	r := gin.Default()
	router.ApiRouter(r)
	return r.Run(fmt.Sprintf("%s%s", ":", global.Config.Http.BindPort))
}
