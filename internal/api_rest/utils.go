package api_rest

import (
	"fmt"
	"github.com/forwork0529/for_open_graph_protokol/pkg/logger"
	"github.com/gin-gonic/gin"
	"strconv"
)

func getIntValFromUrl(ctx *gin.Context, key string, defaultVal int64) (int64, error) {
	val := ctx.Param(key)
	res, err := validateInputValue(val, defaultVal)
	if err != nil {
		logger.Infof("cant get value from key: %v in url string: %v", key, ctx.Request.URL)
	}
	return res, err
}

func getIntValFromUrlParams(ctx *gin.Context, key string, defaultVal int64) (int64, error) {
	val := ctx.Query(key)
	res, err := validateInputValue(val, defaultVal)
	if err != nil {
		logger.Infof("cant get value from key: %v in url string: %v", key, ctx.Request.URL)
	}
	return res, err
}

func validateInputValue(val string, defaultVal int64) (int64, error) {
	var err error
	if len(val) == 0 {
		err = fmt.Errorf("cant find parameter in request url")
		logger.Infof("validateInputValue(): %v", err.Error())
		return defaultVal, err
	}
	output, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		err = fmt.Errorf("cant convert to int64 value from the request  url parameters")
		logger.Infof("validateInputValue(): %v", err.Error())
		return defaultVal, err
	}
	if output < 0 {
		err = fmt.Errorf("value from paramenter < 0, returning default: %v", defaultVal)
		logger.Infof("validateInputValue(): %v", err.Error())
		return defaultVal, err
	}
	return int64(output), nil
}
