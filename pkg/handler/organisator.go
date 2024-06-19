package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	app "github.com/reckycless/event-checker"
)

// @Summary Create Organisator
// @Security ApiKeyAuth
// @Tags organisators
// @Description create organisator
// @ID create-organisator
// @Accept  json
// @Produce  json
// @Param input body app.CreateOrganisatorInput true "organisator input"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/v1/private/organisators [post]
func (h *Handler) createOrganisator(c *gin.Context) {
	valid, _ := validateAdminRole(c)
	if !valid {
		return
	}

	var input app.CreateOrganisatorInput
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

// @Summary Get All Organisators
// @Security ApiKeyAuth
// @Tags organisators
// @Description get all organisators
// @ID get-all-organisators
// @Accept  json
// @Produce  json
// @Success 200 {object} []app.GetOrganisatorOutput
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/v1/public/organisators [get]
func (h *Handler) getAllOrganisators(c *gin.Context) {
	organisators, err := h.services.Organisator.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, organisators)
}

// @Summary Get Organisator By ID
// @Tags organisators
// @Description get organisator by id
// @ID get-organisator-by-id
// @Accept  json
// @Produce  json
// @Param id path int true "organisator id"
// @Success 200 {object} app.GetOrganisatorOutput
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/v1/public/organisators/{id} [get]
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

// @Summary Delete Organisator by ID
// @Security ApiKeyAuth
// @Tags organisators
// @Description delete organisator
// @ID delete-organisator
// @Accept  json
// @Produce  json
// @Param id path int true "organisator id"
// @Success 200 {integer} integer 1
// @Failure 204 {object} errorResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/v1/private/organisators/{id} [delete]
func (h *Handler) deleteOrganisator(c *gin.Context) {
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

// @Summary Update Organisator by ID
// @Security ApiKeyAuth
// @Tags organisators
// @Description update organisator
// @ID update-organisator
// @Accept  json
// @Produce  json
// @Param id path int true "organisator id"
// @Param input body app.UpdateOrganisatorInput true "organisator input"
// @Success 200 {integer} integer 1
// @Failure 204 {object} errorResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/v1/private/organisators/{id} [put]
func (h *Handler) updateOrganisator(c *gin.Context) {
	valid, _ := validateAdminRole(c)
	if !valid {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input app.UpdateOrganisatorInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	rowsAffected, err := h.services.Organisator.Update(id, input)
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
