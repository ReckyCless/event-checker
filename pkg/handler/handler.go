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
		v1 := api.Group("/v1")
		{
			//Public API Routes
			public := v1.Group("/public")
			{
				auth := public.Group("/auth")
				{
					auth.POST("/sign-up", h.signUp)
					auth.POST("/sign-in", h.signIn)
				}

				events := public.Group("/events")
				{
					events.GET("/", h.getAllEvents)
					events.GET("/:id", h.getEventByID)

					//Events Visitors Router Group
					visitors := events.Group(":id/visitors")
					{
						visitors.POST("/", h.createVisitor)
					}
				}

				organisators := public.Group("/organisators")
				{
					organisators.GET("/", h.getAllOrganisators)
					organisators.GET("/:id", h.getOrganisatorByID)
				}

				types := public.Group("/types")
				{
					types.GET("/", h.getAllEventTypes)
					types.GET("/:id", h.getEventTypeByID)
				}

				roles := public.Group("/roles")
				{
					roles.GET("/", h.getAllRoles)
					roles.GET("/:id", h.getRoleByID)
				}
			}

			//Private API Routes
			private := v1.Group("/private", h.userIdentity)
			{
				events := private.Group("/events")
				{
					events.GET("/", h.getAllEventsForUser)
					events.POST("/", h.createEvent)
					events.PUT("/:id", h.updateEvent)
					events.DELETE("/:id", h.deleteEvent)

					visitors := events.Group("/:id/visitors")
					{
						visitors.POST("/", h.createVisitorAsUser)
						visitors.GET("/", h.getAllEventVisitors)
						visitors.GET("/:visitor_id", h.getVisitorByID)
						visitors.PUT("/:visitor_id", h.updateVisitor)
						visitors.DELETE("/:visitor_id", h.deleteVisitor)
					}
				}

				organisators := private.Group("/organisators")
				{
					organisators.POST("/", h.createOrganisator)
					organisators.PUT("/:id", h.updateOrganisator)
					organisators.DELETE("/:id", h.deleteOrganisator)
				}

				types := private.Group("/types")
				{
					types.POST("/", h.createEventType)
					types.PUT("/:id", h.updateEventType)
					types.DELETE("/:id", h.deleteEventType)
				}

				roles := private.Group("/roles")
				{
					roles.PUT("/:id", h.updateRole)
				}

				users := private.Group("/users")
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
		}
	}

	return router
}
