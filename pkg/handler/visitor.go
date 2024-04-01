package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	app "github.com/reckycless/event-checker"
)

// @Summary Create Event Visitor
// @Tags visitors
// @Description create event visitor
// @ID create-event-visitor
// @Accept  json
// @Produce  json
// @Param id path int true "event id"
// @Param input body app.Visitor true "visitor info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/events/{id}/visitors [post]
func (h *Handler) createVisitor(c *gin.Context) {
	eventID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input app.Visitor
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Visitor.Create(eventID, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

// @Summary Create Event Visitor as User
// @Security ApiKeyAuth
// @Tags visitors
// @Description create event visitor as user
// @ID user-create-event-visitor
// @Accept  json
// @Produce  json
// @Param id path int true "event id"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/events/authorised/{id}/visitors [post]
func (h *Handler) createVisitorAsUser(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		return
	}

	eventID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	id, err := h.services.Visitor.CreateAsUser(userID, eventID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

// @Summary Get All Event Visitors
// @Security ApiKeyAuth
// @Tags visitors
// @Description get all event visitors
// @ID get-all-event-visitors
// @Accept  json
// @Produce  json
// @Param id path int true "event id"
// @Success 200 {object} []app.Visitor
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/events/authorised/{id}/visitors/ [get]
func (h *Handler) getAllEventVisitors(c *gin.Context) {
	//TODO: EDIT REPOSITORY TO MANAGE ROLES
	valid, _ := validateManagerAdminRole(c)
	if !valid {
		return
	}

	userID, err := getUserID(c)
	if err != nil {
		return
	}

	eventID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	visitors, err := h.services.Visitor.GetAllEventVisitors(userID, eventID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, visitors)
}

// @Summary Get Event Visitor By ID
// @Security ApiKeyAuth
// @Tags visitors
// @Description get event visitor by id
// @ID get-event-visitor-by-id
// @Accept  json
// @Produce  json
// @Param id path int true "event id"
// @Param visitor_id path int true "visitor id"
// @Success 200 {object} app.Visitor
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/events/authorised/{id}/visitors/{visitor_id} [get]
func (h *Handler) getVisitorByID(c *gin.Context) {
	//TODO: EDIT REPOSITORY TO MANAGE ROLES
	valid, _ := validateManagerAdminRole(c)
	if !valid {
		return
	}

	userID, err := getUserID(c)
	if err != nil {
		return
	}

	visitorID, err := strconv.Atoi(c.Param("visitor_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid visitor_id param")
		return
	}

	visitor, err := h.services.Visitor.GetByID(userID, visitorID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, visitor)
}

// @Summary Delete Event Visitor By ID
// @Security ApiKeyAuth
// @Tags visitors
// @Description delete event visitor
// @ID delete-event-visitor
// @Accept  json
// @Produce  json
// @Param id path int true "event id"
// @Param visitor_id path int true "visitor id"
// @Success 200 {integer} integer 1
// @Failure 204 {object} errorResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/events/authorised/{id}/visitors/{visitor_id} [delete]
func (h *Handler) deleteVisitor(c *gin.Context) {
	//TODO: EDIT REPOSITORY TO MANAGE ROLES
	valid, _ := validateManagerAdminRole(c)
	if !valid {
		return
	}

	userID, err := getUserID(c)
	if err != nil {
		return
	}

	visitorID, err := strconv.Atoi(c.Param("visitor_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid visitor_id param")
		return
	}

	rowsAffected, err := h.services.Visitor.Delete(userID, visitorID)
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

// @Summary Update Event Visitor by ID
// @Security ApiKeyAuth
// @Tags visitors
// @Description update event visitor
// @ID update-event-visitor
// @Accept  json
// @Produce  json
// @Param id path int true "event id"
// @Param visitor_id path int true "visitor id"
// @Param input body app.UpdateVisitorInput true "visitor info"
// @Success 200 {integer} integer 1
// @Failure 204 {object} errorResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/events/authorised/{id}/visitors/{visitor_id} [put]
func (h *Handler) updateVisitor(c *gin.Context) {
	//TODO: EDIT REPOSITORY TO MANAGE ROLES
	valid, _ := validateManagerAdminRole(c)
	if !valid {
		return
	}

	userID, err := getUserID(c)
	if err != nil {
		return
	}

	visitorID, err := strconv.Atoi(c.Param("visitor_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid visitor id param")
		return
	}

	var input app.UpdateVisitorInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	rowsAffected, err := h.services.Visitor.Update(userID, visitorID, input)
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
