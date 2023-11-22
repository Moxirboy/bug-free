package handler

import (
	"bug-free/internal/model"
	"bug-free/internal/response"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createUser(ctx *gin.Context) {
	var user model.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		response.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	fmt.Print("here is it :"+user.Name)
	id, err := h.service.Create(user)
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
func (h *Handler) main(ctx *gin.Context) {
	ctx.String(200, "Hello, World!")
}
func (h *Handler) InitRouter() *gin.Engine {
	router := gin.New()
	crud := router.Group("/")
	{
		crud.GET("/" , h.main)
		crud.POST("/create", h.createUser)
		crud.POST("/read", h.GetUser)
		crud.POST("/update", h.UpdateUser)
		crud.POST("/delete", h.DeleteUser)

	}

	return router
}
