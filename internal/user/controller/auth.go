package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/trivery-id/skadi/internal/user/services"
	"gitlab.com/trivery-id/skadi/utils/errors"
	"gitlab.com/trivery-id/skadi/utils/metadata"
	writer "gitlab.com/trivery-id/skadi/utils/response-writer"
)

func (ctrl *AuthController) Login(c *gin.Context) {
	var in services.LoginInput
	if err := c.ShouldBindJSON(&in); err != nil {
		writer.WriteFailResponseFromError(c, errors.NewBadRequestError(writer.ErrMsgUnableToBindJSON, err))
		return
	}

	resp, err := ctrl.UserService.Login(c.Request.Context(), in)
	if err != nil {
		writer.WriteFailResponseFromError(c, err)
		return
	}

	writer.WriteSuccessResponse(c, http.StatusOK, resp)
}

func (ctrl *AuthController) RefreshToken(c *gin.Context) {
	var in services.RefreshTokenInput
	if err := c.ShouldBindJSON(&in); err != nil {
		writer.WriteFailResponseFromError(c, errors.NewBadRequestError(writer.ErrMsgUnableToBindJSON, err))
		return
	}

	resp, err := ctrl.UserService.RefreshToken(c.Request.Context(), in)
	if err != nil {
		writer.WriteFailResponseFromError(c, err)
		return
	}

	writer.WriteSuccessResponse(c, http.StatusOK, resp)
}

func (*AuthController) Test(c *gin.Context) {
	user := metadata.GetUserFromContext(c.Request.Context())
	writer.WriteSuccessResponse(c, http.StatusOK, user)
}
