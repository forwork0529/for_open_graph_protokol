package api_rest

import (
	"github.com/forwork0529/for_open_graph_protokol/pkg/logger"
	l "github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
)

type Server struct {
}

func NewServer() *Server {
	return &Server{}
}

func RunGin(s *Server, endpoint string) {
	// инициализируем gin, намерено не делаю его полем сервера:
	e := gin.New()
	e.Use(l.SetLogger())
	e.Use(gin.Recovery())
	// e.Use(CORSMiddleware()) not need

	e.GET("/health", s.health)

	api := e.Group("/api/v1")
	{
		group := api.Group("/static")
		{
			group.GET("/:element_type/:element_id", s.getElement)
		}
	}

	// запуск API
	go func() {
		err := e.Run(endpoint)
		if err != nil {
			logger.Errorf("RunGin(): %v", err)
		}
	}()

	logger.Infof("Gin HTTP server successfully started..")

}
