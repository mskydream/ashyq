package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mskydream/qr-code/cmd/model"
)

func (h *Handler) getAllRealEstate(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, model.Response{IsSuccess: false, Message: "Введенные данные некорректны", Data: err})
		return
	}
	realEstate, err := h.service.GetRealEstate(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, model.Response{IsSuccess: false, Message: "Ошибка в сервере", Data: err})
		return
	}

	c.JSON(http.StatusOK, model.Response{IsSuccess: true, Message: "Успешно получено", Data: realEstate})

}
