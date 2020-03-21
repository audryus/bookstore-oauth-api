package http

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gitlab.com/aubayaml/aubayaml-go/bookstore/oauth-api/domain/accesstoken"
	"gitlab.com/aubayaml/aubayaml-go/bookstore/utils-go/errors"
)

//AccessTokenHandler token
type AccessTokenHandler interface {
	GetByID(c *gin.Context)
	Create(c *gin.Context)
	UpdateExpirationTime(c *gin.Context)
}
type accessTokenHandler struct {
	service accesstoken.Service
}

//NewHandler with the given service
func NewHandler(s accesstoken.Service) AccessTokenHandler {
	return &accessTokenHandler{
		service: s,
	}
}

func (h *accessTokenHandler) GetByID(c *gin.Context) {
	token, err := h.service.GetByID(strings.TrimSpace(c.Param("token_id")))
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, token)
}

func (h *accessTokenHandler) Create(c *gin.Context) {
	var lr accesstoken.LoginRequest
	if err := c.ShouldBindJSON(&lr); err != nil {
		restErr := errors.BadRequestError("invalid json body", err)
		c.JSON(restErr.Status(), restErr)
		return
	}

	at, err := h.service.Create(lr)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusCreated, at)
}

func (h *accessTokenHandler) UpdateExpirationTime(c *gin.Context) {
}
