package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	app "github.com/reckycless/event-checker"
)

func (h *Handler) updateUserRole(c *gin.Context) {
	//TODO: EDIT REPOSITORY TO MANAGE ROLES
	valid, _ := validateAdminRole(c)
	if !valid {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input app.UserRoleInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	rowsAffected, err := h.services.User.UpdateUserRole(id, input)
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

func (h *Handler) updateUserOrganisator(c *gin.Context) {
	//TODO: EDIT REPOSITORY TO MANAGE ROLES
	valid, _ := validateAdminRole(c)
	if !valid {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input app.UserOrganisatorInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	rowsAffected, err := h.services.User.UpdateUserOrganisator(id, input)
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
