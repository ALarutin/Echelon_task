package handlers

import (
	log "github.com/ALarutin/Echelon_task/logger"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

var logger *log.Logger

func init() {
	logger = log.GetLogger("Main")
	logger.SetLogLevel(logrus.TraceLevel)
}

func CheckText(c *gin.Context) {
	var req Request
	err := c.BindJSON(&req)
	if err != nil {
		logger.Error("Failed to bind json with error: %v", err.Error())
	}
	logger.Debugf("Sites: %v", req.Sites)
	logger.Debugf("Search text: %v", req.SearchText)

	c.String(http.StatusOK, "Hello %v, %s", req.Sites, req.SearchText)
}
