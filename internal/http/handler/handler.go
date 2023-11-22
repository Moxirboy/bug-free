package handler

import (
	"github.com/gin-gonic/gin"
	"golang-project-template/internal/model"
	"golang-project-template/internal/response"
	"net/http"
)

func (h *Handler) createUser(ctx *gin.Context) {
	var User model.User
	if err := ctx.BindJSON(User); err != nil {
		response.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.service.Create(User)
	if err != nil {
		response.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{"id": id})

}
func (h *Handler) GetUser(ctx *gin.Context) {
	var Users = h.service.Read()
	_ = Users
}
func (h *Handler) UpdateUser(ctx *gin.Context) {
	h.service.Update()
}
func (h *Handler) DeleteUser(ctx *gin.Context) {
	h.service.Delete()
}
func (h *Handler) InitRouter() *gin.Engine {
	router := gin.New()
	crud := router.Group("/")
	{
		crud.POST("/Create", h.createUser)
		crud.POST("/Read", h.GetUser)
		crud.POST("/Update", h.UpdateUser)
		crud.POST("/Delete", h.DeleteUser)

	}

	return router
}
