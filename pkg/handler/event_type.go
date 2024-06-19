package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	app "github.com/reckycless/event-checker"
)

// @Summary Create Event Type
// @Security ApiKeyAuth
// @Tags types
// @Description create event type
// @ID create-event-type
// @Accept  json
// @Produce  json
// @Param input body app.CreateEventTypeInput true "event type input"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/v1/private/types [post]
func (h *Handler) createEventType(c *gin.Context) {
	valid, _ := validateAdminRole(c)
	if !valid {
		return
	}

	var input app.CreateEventTypeInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.EventType.Create(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

// @Summary Get All Events Types
// @Tags types
// @Description get all events created by user
// @ID get-all-events-types
// @Accept  json
// @Produce  json
// @Success 200 {object} []app.GetEventTypeOutput
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/v1/public/types [get]
func (h *Handler) getAllEventTypes(c *gin.Context) {
	eventTypes, err := h.services.EventType.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, eventTypes)
}

// @Summary Get Event Type By ID
// @Tags types
// @Description get event type by id
// @ID get-event-type-by-id
// @Accept  json
// @Produce  json
// @Param id path int true "event type id"
// @Success 200 {object} app.GetEventTypeOutput
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/v1/public/types/{id} [get]
func (h *Handler) getEventTypeByID(c *gin.Context) {
	eventTypeID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	eventType, err := h.services.EventType.GetByID(eventTypeID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, eventType)
}

// @Summary Delete Event Type by ID
// @Security ApiKeyAuth
// @Tags types
// @Description delete event type
// @ID delete-type
// @Accept  json
// @Produce  json
// @Param id path int true "event type id"
// @Success 200 {integer} integer 1
// @Failure 204 {object} errorResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/v1/private/types/{id} [delete]
func (h *Handler) deleteEventType(c *gin.Context) {
	valid, _ := validateAdminRole(c)
	if !valid {
		return
	}

	eventTypeID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	rowsAffected, err := h.services.EventType.Delete(eventTypeID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected < 1 {
		newErrorResponse(c, http.StatusNoContent, "no rows in result set")
		return
	}

	c.JSON(http.StatusOK, statusResponse{Status: "ok"})
}

// @Summary Update Event Type by ID
// @Security ApiKeyAuth
// @Tags types
// @Description update event type
// @ID update-type
// @Accept  json
// @Produce  json
// @Param id path int true "event type id"
// @Param input body app.UpdateEventTypeInput true "event type input"
// @Success 200 {integer} integer 1
// @Failure 204 {object} errorResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/v1/private/types/{id} [put]
func (h *Handler) updateEventType(c *gin.Context) {
	valid, _ := validateAdminRole(c)
	if !valid {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input app.UpdateEventTypeInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	rowsAffected, err := h.services.EventType.Update(id, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected < 1 {
		newErrorResponse(c, http.StatusNoContent, "no rows in result set")
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
