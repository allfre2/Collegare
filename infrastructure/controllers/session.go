package controllers

import (
	"collegare/application/services"
	"collegare/domain"

	"github.com/gin-gonic/gin"
)

type SessionController struct {
	sessionSvc services.SessionService
}

func NewSessionController(sessionSvc services.SessionService) *SessionController {
	return &SessionController{
		sessionSvc: sessionSvc,
	}
}

func (ctrl *SessionController) Session(ctx *gin.Context) {
	ctrl.sessionSvc.Connect(&domain.SessionRequest{Context: ctx})
}
