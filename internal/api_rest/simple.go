package api_rest

import (
	"github.com/forwork0529/for_open_graph_protokol/pkg/models"
	"github.com/gin-gonic/gin"

	"net/http"
)

func (s *Server) health(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, models.StatusResponse{Status: "OK", Message: "health all right"})
}
