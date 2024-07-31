package helpers

import (
	"github.com/malpania/beerproj/pkg/config"
	"net/http"
	"runtime/debug"
)

var appConfig *config.AppConfig

func NewHelpers(config *config.AppConfig) {
	appConfig = config
}

func ClientError(w http.ResponseWriter, status int) {
	appConfig.InfoLog.Println("Client Error:", status)
	http.Error(w, http.StatusText(status), status)
}

func ServerError(w http.ResponseWriter, err error) {
	appConfig.ErrorLog.Println("Unexpected error : ", err.Error(), debug.Stack())
	http.Error(w, err.Error(), http.StatusInternalServerError)
}
