package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mskydream/ashyq/model"
)

func (h *Handler) getVisitByQrCode(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, model.Response{IsSuccess: false, Message: "Необходима авторизация", Data: err})
		return
	}

	qrId := c.Param("id")
	status, err := h.service.GetStatus(userId, qrId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, model.Response{IsSuccess: false, Message: "Ошибка в сервере", Data: err})
		return
	}

	c.JSON(http.StatusOK, model.Response{IsSuccess: true, Message: "Успешно создано", Data: status})
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
