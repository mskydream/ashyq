package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mskydream/ashyq/model"
)

func (h *Handler) createRealEstate(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, model.Response{IsSuccess: false, Message: "Необходима авторизация", Data: err})
		return
	}

	var realEstate model.RealEstate
	if err := c.ShouldBindJSON(&realEstate); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, model.Response{IsSuccess: false, Message: "Введенные данные некорректны", Data: err})
		return
	}

	if err := h.service.CheckAddress(realEstate.Address); err == nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, model.Response{IsSuccess: false, Message: "Адрес уже добавлен", Data: ""})
		return
	}

	id, err := h.service.Create(userId, &realEstate)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, model.Response{IsSuccess: false, Message: "Ошибка в сервере", Data: err})
		return
	}

	c.JSON(http.StatusCreated, model.Response{IsSuccess: true, Message: "Объект создан", Data: id})
}

func (h *Handler) getAllRealEstates(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, model.Response{IsSuccess: false, Message: "Необходима авторизация", Data: err})
		return
	}

	realEstates, err := h.service.GetAll(userId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, model.Response{IsSuccess: false, Message: "Ошибка в сервере", Data: err})
		return
	}

	c.JSON(http.StatusOK, model.Response{IsSuccess: true, Message: "Объекты получены", Data: realEstates})
}

func (h *Handler) getRealEstate(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, model.Response{IsSuccess: false, Message: "Необходима авторизация", Data: err})
		return
	}

	id := c.Param("id")
	realEstate, err := h.service.Get(userId, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, model.Response{IsSuccess: false, Message: "Ошибка в сервере", Data: err})
		return
	}

	c.JSON(http.StatusOK, model.Response{IsSuccess: true, Message: "Объект получен", Data: realEstate})
}

func (h *Handler) deleteRealEstate(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, model.Response{IsSuccess: false, Message: "Необходима авторизация", Data: err})
		return
	}

	id := c.Param("id")
	err = h.service.Delete(userId, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, model.Response{IsSuccess: false, Message: "Ошибка в сервере", Data: err})
		return
	}

	c.JSON(http.StatusOK, model.Response{IsSuccess: true, Message: "Объект удален", Data: "deleted"})
}

func (h *Handler) updateRealEstate(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, model.Response{IsSuccess: false, Message: "Необходима авторизация", Data: err})
		return
	}

	id := c.Param("id")
	var realEstate model.RealEstate
	if err := c.ShouldBindJSON(&realEstate); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, model.Response{IsSuccess: false, Message: "Введенные данные некорректны", Data: err})
		return
	}

	err = h.service.Update(userId, id, &realEstate)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, model.Response{IsSuccess: false, Message: "Ошибка в сервере", Data: err})
		return
	}

	c.JSON(http.StatusOK, model.Response{IsSuccess: true, Message: "Объект обновлен", Data: "sucess"})
}
