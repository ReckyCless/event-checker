package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	app "github.com/reckycless/event-checker"
)

// @Summary Create Event
// @Security ApiKeyAuth
// @Tags events
// @Description create event
// @ID create-event
// @Accept  json
// @Produce  json
// @Param input body app.CreateEventInput true "event input"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/v1/private/events [post]
func (h *Handler) createEvent(c *gin.Context) {
	valid, _ := validateManagerAdminRole(c)
	if !valid {
		return
	}

	userID, err := getUserID(c)
	if err != nil {
		return
	}

	var input app.CreateEventInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Event.Create(userID, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

// @Summary Get All Events by User
// @Security ApiKeyAuth
// @Tags events
// @Description get all events created by user
// @ID get-all-user-events
// @Accept  json
// @Produce  json
// @Success 200 {object} []app.GetEventOutput
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/v1/private/events [get]
func (h *Handler) getAllEventsForUser(c *gin.Context) {
	valid, _ := validateManagerAdminRole(c)
	if !valid {
		return
	}

	userID, err := getUserID(c)
	if err != nil {
		return
	}

	events, err := h.services.Event.GetAllForUser(userID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, events)
}

// @Summary Get All Events
// @Tags events
// @Description get all events
// @ID get-all-events
// @Accept  json
// @Produce  json
// @Success 200 {object} []app.GetEventOutput
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/v1/public/events [get]
func (h *Handler) getAllEvents(c *gin.Context) {
	events, err := h.services.Event.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, events)
}

// @Summary Get Event By ID
// @Tags events
// @Description get event by id
// @ID get-event-by-id
// @Accept  json
// @Produce  json
// @Param id path int true "event id"
// @Success 200 {object} app.GetEventOutput
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/v1/public/events/{id} [get]
func (h *Handler) getEventByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	event, err := h.services.Event.GetByID(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, event)
}

// @Summary Delete Event by ID
// @Security ApiKeyAuth
// @Tags events
// @Description delete event
// @ID delete-event
// @Accept  json
// @Produce  json
// @Param id path int true "event id"
// @Success 200 {integer} integer 1
// @Failure 204 {object} errorResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/v1/private/events/{id} [delete]
func (h *Handler) deleteEvent(c *gin.Context) {
	valid, _ := validateManagerAdminRole(c)
	if !valid {
		return
	}

	userID, err := getUserID(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	rowsAffected, err := h.services.Event.Delete(userID, id)
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

// @Summary Update Event by ID
// @Security ApiKeyAuth
// @Tags events
// @Description update event
// @ID update-event
// @Accept  json
// @Produce  json
// @Param id path int true "event id"
// @Param input body app.UpdateEventInput true "event input"
// @Success 200 {integer} integer 1
// @Failure 204 {object} errorResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/v1/private/events/{id} [put]
func (h *Handler) updateEvent(c *gin.Context) {
	valid, _ := validateManagerAdminRole(c)
	if !valid {
		return
	}

	userID, err := getUserID(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input app.UpdateEventInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	rowsAffected, err := h.services.Event.Update(userID, id, input)
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
