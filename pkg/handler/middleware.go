package handler

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type UserCtx struct {
	UserID string
	RoleID string
}

// Constants for headers
var userCtx = UserCtx{"userID", "roleID"}

const (
	authorizationHeader = "Authorization"
	adminRoleID         = 1
	ManagerRoleID       = 2
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		newErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
		return
	}

	userID, roleID, err := h.services.Authorization.ParseToken(headerParts[1])
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set(userCtx.UserID, userID)
	c.Set(userCtx.RoleID, roleID)
}

func getUserID(c *gin.Context) (int, error) {
	id, ok := c.Get(userCtx.UserID)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user id not found")
		return 0, errors.New("user id not found")
	}

	idInt, ok := id.(int)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user id is of invalid type")
		return 0, errors.New("user id is of invalid type")
	}

	return idInt, nil
}

func getRoleID(c *gin.Context) (int, error) {
	id, ok := c.Get(userCtx.RoleID)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "role id not found")
		return 0, errors.New("role id not found")
	}

	idInt, ok := id.(int)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "role id is of invalid type")
		return 0, errors.New("role id is of invalid type")
	}

	return idInt, nil
}

func validateAdminRole(c *gin.Context) (bool, error) {
	roleID, err := getRoleID(c)
	if err != nil {
		return false, err
	}

	if roleID != adminRoleID {
		newErrorResponse(c, http.StatusForbidden, "insufficient role")
		return false, errors.New("insufficient role")
	}

	return true, nil
}

func validateManagerAdminRole(c *gin.Context) (bool, error) {
	roleID, err := getRoleID(c)
	if err != nil {
		return false, err
	}

	if roleID != adminRoleID && roleID != ManagerRoleID {
		newErrorResponse(c, http.StatusForbidden, "insufficient role")
		return false, errors.New("insufficient role")
	}

	return true, nil
}
