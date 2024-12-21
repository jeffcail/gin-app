package boot

import (
	_uber "github.com/mazezen/ginframe/server-common/pkg/uber"
	"github.com/mazezen/ginframe/server-user/ulogger"
)

// InitLog init log
func InitLog(filepath string) {
	logger := _uber.InitLogger(filepath)
	ulogger.SetLogger(logger)

	ulogger.UserLogger.Info("🚀🚀🚀🚀🚀🚀 logger success...🚀🚀🚀🚀🚀🚀")
}
