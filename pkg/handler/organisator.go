package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	app "github.com/reckycless/event-checker"
)

func (h *Handler) createOrganisator(c *gin.Context) {
	//TODO: EDIT REPOSITORY TO MANAGE ROLES
	valid, _ := validateAdminRole(c)
	if !valid {
		return
	}

	var input app.Organisator
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Organisator.Create(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getAllOrganisators(c *gin.Context) {
	visitors, err := h.services.Organisator.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, visitors)
}

func (h *Handler) getOrganisatorByID(c *gin.Context) {
	organisatorID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	organisator, err := h.services.Organisator.GetByID(organisatorID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, organisator)
}

func (h *Handler) deleteOrganisator(c *gin.Context) {
	//TODO: EDIT REPOSITORY TO MANAGE ROLES
	valid, _ := validateAdminRole(c)
	if !valid {
		return
	}

	organisatorID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	rowsAffected, err := h.services.Organisator.Delete(organisatorID)
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

func (h *Handler) updateOrganisator(c *gin.Context) {
	//TODO: EDIT REPOSITORY TO MANAGE ROLES
	valid, _ := validateAdminRole(c)
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
