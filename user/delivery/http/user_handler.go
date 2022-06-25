package http

import (
	"go-redis/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userUsecase domain.UserUsecase
}

func NewUserHandler(r *gin.RouterGroup, userUsecase domain.UserUsecase) {
	handler := &UserHandler{userUsecase}
	api := r.Group("/user")
	{
		api.POST("/seed", handler.Seed)
		api.GET("/:id/withoutredis", handler.GetByIdWithoutRedis)
		api.GET("/:id/withredis", handler.GetByIdWithRedis)
	}
}

func (h *UserHandler) Seed(c *gin.Context) {
	if err := h.userUsecase.Seed(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"msg": "successfully seeding database"})
}

func (h *UserHandler) GetByIdWithoutRedis(c *gin.Context) {
	var uri domain.UserIdUriBinding
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	user, err := h.userUsecase.GetByIdWithoutRedis(uri.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) GetByIdWithRedis(c *gin.Context) {
	var uri domain.UserIdUriBinding
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	user, err := h.userUsecase.GetByIdWithRedis(uri.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}
