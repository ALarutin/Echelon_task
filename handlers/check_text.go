package handlers

import (
	log "github.com/ALarutin/Echelon_task/logger"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"regexp"
	"strings"
)

var logger *log.Logger

func init() {
	logger = log.GetLogger("Handlers")
	logger.SetLogLevel(logrus.ErrorLevel)
}

func CheckText(c *gin.Context) {
	var req Request
	err := c.BindJSON(&req)
	if err != nil {
		logger.Error("Failed to bind json with error: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	logger.Debugf("Sites: %v", req.Sites)
	logger.Debugf("Search text: %s", req.SearchText)

	req.SearchText = strings.ToLower(req.SearchText)
	logger.Debugf("Search text after lower: %s", req.SearchText)

	var res Response
	for _, s := range req.Sites {
		matched, err := regexp.Match(req.SearchText, []byte(s))
		if err != nil {
			logger.Error("Failed to matched with error: %s", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if matched {
			res.FoundAtSites = append(res.FoundAtSites, s)
			logger.Debugf("Matched site: %s", s)
		}
	}
	logger.Debugf("Matched sites: %v", res.FoundAtSites)

	c.JSON(http.StatusOK, res)
}
