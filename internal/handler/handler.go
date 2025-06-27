package handler

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine{
	router := gin.New()

	base := router.Group("/w")
	{
		base.GET("/s", handler.HandlerWebsocket)
	}
	return router
}

func (h *Handler) HandlerWebsocket(c *gin.Context) {
	
}