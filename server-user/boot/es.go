package boot

import (
	"github.com/mazezen/ginframe/server-common/driver"
	"github.com/mazezen/ginframe/server-user/core"
	"github.com/mazezen/ginframe/server-user/ulogger"
	"log"
)

// InitEs init elastic
func InitEs(url string) {
	es, err := driver.InitEs(url)
	if err != nil {
		log.Fatalln(err)
	}
	core.SetEsClient(es)

	ulogger.UserLogger.Info("🚀🚀🚀🚀🚀🚀ES success...🚀🚀🚀🚀🚀🚀")
}
