package handlers

import (
	"net/http"
	"service-golang/models"
	"service-golang/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct{
	UserService *services.UserService
}

func NewUserHandler(service *services.UserService) *UserHandler{
	return &UserHandler{UserService: service}
}

func (h *UserHandler) CreateUser(c *gin.Context){
	var user models.User
	if err := c.ShouldBind(&user); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.UserService.CreateUser(&user); err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully create user"})
}

func (h *UserHandler) GetAllUser(c *gin.Context){
	users, err := h.UserService.GetAllUser()
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"datas": users, "message": "OK"})
}

func (h *UserHandler) GetUserByID(c *gin.Context){
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10,32)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.UserService.GetUserByID(uint(id))
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"datas": user, "message": "OK"})
}

func (h *UserHandler) UpdateUser(c *gin.Context){
	var user models.User
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10,32)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.ShouldBind(&user); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.UserService.UpdateUser(&user, uint(id)); err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"messsage": "Successfully Update User"})
}

func (h *UserHandler) DeleteUser(c *gin.Context){
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10,32)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.UserService.DeleteUser(uint(id)); err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully delete User"})
}