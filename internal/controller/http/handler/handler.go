package handler

import (
	"bug-free/internal/model"
	"bug-free/internal/response"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @Summary   	Send code to telegram group
// @Description This api can send code
// @Tags     	Send
// @Accept     	json
// @Produce   	json
// @Param     	body body model.UserCreate true "User"
// @Failure 400 string Error response
// @Router /v1/sign-up-verify-email [get]
func (h *Handler) createUser(ctx *gin.Context) {
	var user model.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		response.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	fmt.Print("here is it :" + user.Name)
	id, err := h.service.Create(user)
	if err != nil {
		response.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{"id": id})

}
func (h *Handler) GetUser(ctx *gin.Context) {
	var userID model.UserGetBYId
	if err := ctx.ShouldBindJSON(&userID); err != nil {
		response.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	name, err := h.service.Read(userID)
	if err != nil {
		response.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{"name": name})
}
func (h *Handler) UpdateUser(ctx *gin.Context) {
	var User model.User
	if err := ctx.ShouldBindJSON(&User); err != nil {
		response.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.service.Update(User)

	if err != nil {
		response.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{"id": id})

}
func (h *Handler) DeleteUser(ctx *gin.Context) {
	var user model.User
	if err := ctx.ShouldBind(&user); err != nil {
		response.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.service.Delete(user)
	if err != nil {
		response.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{"id": id})
}
func (h *Handler) main(ctx *gin.Context) {
	ctx.String(200, "Hello, World!")
}
func (h *Handler) InitRouter() *gin.Engine {
	router := gin.New()
	crud := router.Group("/")
	url := ginSwagger.URL("swagger/doc.json")
	{
		crud.GET("/", h.main)
		crud.POST("/create", h.createUser)
		crud.POST("/read", h.GetUser)
		crud.POST("/update", h.UpdateUser)
		crud.POST("/delete", h.DeleteUser)
		crud.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	}

	return router
}
