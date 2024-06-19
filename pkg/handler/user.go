package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	app "github.com/reckycless/event-checker"
)

// @Summary User Role Modification
// @Security ApiKeyAuth
// @Tags users
// @Description update user role
// @ID update-user-role
// @Accept  json
// @Produce  json
// @Param id path int true "user id"
// @Param input body app.UserRoleInput true "role input"
// @Success 200 {integer} integer 1
// @Failure 204 {object} errorResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/v1/private/users/{id}/roles [put]
func (h *Handler) updateUserRole(c *gin.Context) {
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

// @Summary User Affilation Modification
// @Security ApiKeyAuth
// @Tags users
// @Description update user organisator
// @ID update-user-organisator
// @Accept  json
// @Produce  json
// @Param id path int true "user id"
// @Param input body app.UserOrganisatorInput true "role input"
// @Success 200 {integer} integer 1
// @Failure 204 {object} errorResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/v1/private/users/{id}/organisators [put]
func (h *Handler) updateUserOrganisator(c *gin.Context) {
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
