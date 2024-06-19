package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	app "github.com/reckycless/event-checker"
)

// @Summary Get All Roles
// @Tags roles
// @Description get all roles
// @ID get-all-roles
// @Accept  json
// @Produce  json
// @Success 200 {object} []app.Role
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/v1/public/roles [get]
func (h *Handler) getAllRoles(c *gin.Context) {
	roles, err := h.services.Role.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, roles)
}

// @Summary Get Role By ID
// @Tags roles
// @Description get role by id
// @ID get-role-by-id
// @Accept  json
// @Produce  json
// @Param id path int true "role id"
// @Success 200 {object} app.Role
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/v1/public/roles/{id} [get]
func (h *Handler) getRoleByID(c *gin.Context) {
	roleID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	role, err := h.services.Role.GetByID(roleID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, role)
}

// @Summary Update Role by ID
// @Security ApiKeyAuth
// @Tags roles
// @Description update role
// @ID update-role
// @Accept  json
// @Produce  json
// @Param id path int true "role id"
// @Param input body app.UpdateRoleInput true "role input"
// @Success 200 {integer} integer 1
// @Failure 204 {object} errorResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/v1/private/roles/{id} [put]
func (h *Handler) updateRole(c *gin.Context) {
	valid, _ := validateAdminRole(c)
	if !valid {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input app.UpdateRoleInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	rowsAffected, err := h.services.Role.Update(id, input)
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
