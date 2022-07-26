package delivery

import (
	"github.com/jackthepanda96/learn-deploy/domain"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	useCase domain.UserUseCase
}

func New(e *echo.Echo, uc domain.UserUseCase) {

}
