package log

import (
	"strings"

	appLogger "go-hexagonal/utils/logger"

	"bitbucket.org/klopos/majoo-logger/log"
	"github.com/labstack/echo/v4"
)

// Constant variable in API log. 2nd level logging we put into Event.
// Example : log.Level_1().Msg(Level2)

// APILogHandler : handle something who need to do
func APILogHandler(c echo.Context, req, res []byte) {
	appLogger.Logger.Info("log echo:  global log request response",
		log.WithRequired(log.Required{
			RequestParams:   string(req),
			ResponseMessage: string(res),
			ErrorCode:       "",
		}),
		log.WithData(map[string]interface{}{
			"context": c,
		}))
}

// APILogSkipper : rules for APILogHandler
func APILogSkipper(c echo.Context) bool {
	// bool, is this url request include "/api"?
	return strings.Contains(c.Request().RequestURI, "/api")
}
