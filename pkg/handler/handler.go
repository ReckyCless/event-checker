package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/reckycless/event-checker/pkg/service"

	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware

	_ "github.com/reckycless/event-checker/docs"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("/api")
	{
		events := api.Group("/events")
		{
			events.GET("/", h.getAllEvents)
			events.GET("/:id", h.getEventByID)
			authorised := events.Group("/authorised", h.userIdentity)
			{
				authorised.GET("/", h.getAllEventsForUser)
				authorised.POST("/", h.createEvent)
				authorised.PUT("/:id", h.updateEvent)
				authorised.DELETE("/:id", h.deleteEvent)

				visitors := authorised.Group("/:id/visitors")
				{
					visitors.POST("/", h.createVisitorAsUser)
					visitors.GET("/", h.getAllEventVisitors)
					visitors.GET("/:visitor_id", h.getVisitorByID)
					visitors.PUT("/:visitor_id", h.updateVisitor)
					visitors.DELETE("/:visitor_id", h.deleteVisitor)
				}
			}
			visitors := events.Group(":id/visitors")
			{
				visitors.POST("/", h.createVisitor)
			}
		}

		organisators := api.Group("/organisators")
		{
			organisators.GET("/", h.getAllOrganisators)
			organisators.GET("/:id", h.getOrganisatorByID)
			authorised := organisators.Group("/authorised", h.userIdentity)
			{
				authorised.POST("/", h.createOrganisator)
				authorised.PUT("/:id", h.updateOrganisator)
				authorised.DELETE("/:id", h.deleteOrganisator)
			}
		}

		types := api.Group("/types")
		{
			types.GET("/", h.getAllEventTypes)
			types.GET("/:id", h.getEventTypeByID)
			authorised := types.Group("/authorised", h.userIdentity)
			{
				authorised.POST("/", h.createEventType)
				authorised.PUT("/:id", h.updateEventType)
				authorised.DELETE("/:id", h.deleteEventType)
			}
		}

		users := api.Group("/users")
		{
			roles := users.Group("/:id/roles")
			{
				roles.PUT("/", h.updateUserRole)
			}
			organisators := users.Group("/:id/organisators")
			{
				organisators.PUT("/", h.updateUserOrganisator)
			}
		}
	}

	return router
}
