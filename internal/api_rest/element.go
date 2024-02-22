package api_rest

import (
	"fmt"
	"github.com/forwork0529/for_open_graph_protokol/pkg/logger"
	"github.com/gin-gonic/gin"
)

func (s *Server) getElement(ctx *gin.Context) {
	fileId, err := getIntValFromUrl(ctx, "element_type", 0)
	if err != nil {
		logger.Errorf("(s *Server) getElement(): %v", err)
	}
	ctx.File(fmt.Sprintf("static/%d.html", fileId))

	//ctx.JSON(http.StatusOK, models.StatusResponse{Status: "OK", Message: "health all right"})
}
