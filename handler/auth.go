package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mskydream/ashyq/model"
)

// @Summary SignUp
// @Tags auth
// @Description create account
// @ID create-account
// @Accept  json
// @Produce  json
// @Param input body model.User true "account info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} model.Response
// @Failure 500 {object} model.Response
// @Failure default {object} model.Response
// @Router /auth/sign-up [post]
func (h *Handler) signUp(c *gin.Context) {
	var input model.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, model.Response{IsSuccess: false, Message: "Введенные данные некорректны", Data: err})
		return
	}

	status, err := h.service.CreateUser(&input)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, model.Response{IsSuccess: false, Message: "Ошибка в сервере", Data: err})
		return
	}

	c.JSON(http.StatusOK, model.Response{IsSuccess: true, Message: "Успешно создано", Data: status})
}

// @Summary SignIn
// @Tags auth
// @Description login
// @ID login
// @Accept  json
// @Produce  json
// @Param input body model.SignInInput true "credentials"
// @Success 200 {string} string "token"
// @Failure 400,404 {object} model.Response
// @Failure 500 {object} model.Response
// @Failure default {object} model.Response
// @Router /auth/sign-in [post]
func (h *Handler) signIn(c *gin.Context) {
	var input model.SignInInput

	if err := c.BindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, model.Response{IsSuccess: false, Message: "Введенные данные некорректны", Data: err})
		return
	}

	token, err := h.service.GenerateToken(input.IIN, input.Password)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, model.Response{IsSuccess: false, Message: "Веденные данные не корректны.", Data: err})
		return
	}

	c.JSON(http.StatusOK, model.Response{IsSuccess: true, Message: "Вход выполнено", Data: token})
}
