package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mskydream/ashyq/cmd/model"
)

func (h *Handler) createVisit(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, model.Response{IsSuccess: false, Message: "Необходима авторизация", Data: err})
		return
	}

	var input model.Visit
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, model.Response{IsSuccess: false, Message: "Введенные данные некорректны", Data: err})
		return
	}

	id, err := h.service.CreateVisit(userId, &input)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, model.Response{IsSuccess: false, Message: "Ошибка в сервере в созданий посещения", Data: err})
		return
	}

	c.JSON(http.StatusOK, model.Response{IsSuccess: true, Message: "Успешно создано", Data: id})
}

func (h *Handler) getAllVisits(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, model.Response{IsSuccess: false, Message: "Необходима авторизация", Data: err})
		return
	}

	visits, err := h.service.GetVisits(userId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, model.Response{IsSuccess: false, Message: "Ошибка в сервере", Data: err})
		return
	}

	c.JSON(http.StatusOK, model.Response{IsSuccess: true, Message: "Успешно получено", Data: visits})
}

func (h *Handler) getVisit(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, model.Response{IsSuccess: false, Message: "Необходима авторизация", Data: err})
		return
	}

	id := c.Param("id")
	visit, err := h.service.GetVisit(userId, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, model.Response{IsSuccess: false, Message: "Ошибка в сервере", Data: err})
		return
	}

	c.JSON(http.StatusOK, model.Response{IsSuccess: true, Message: "Успешно получено", Data: visit})
}
