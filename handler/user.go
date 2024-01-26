package handler

import (
	"github.com/boularbahsmail/go-templ/models"
	"github.com/boularbahsmail/go-templ/view/user"
	"github.com/labstack/echo/v4"
)

type UserHandler struct{}

func (handler UserHandler) HandlerUserShow(context echo.Context) error {
	newUser := models.User{
		Name:     "Ismail Boularbah",
		Email:    "boularbahismail01@gmail.com",
		Provider: "Google",
	}
	return render(context, user.Show(newUser))
}
