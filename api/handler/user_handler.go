package handler

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/vcnt72/try-golang-database-lib/api/presenter"
	"github.com/vcnt72/try-golang-database-lib/api/request"
	"github.com/vcnt72/try-golang-database-lib/entity"
	"github.com/vcnt72/try-golang-database-lib/usecase/user"
)

func register(userService user.Usecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)

		defer cancel()
		var registerRequest request.RegisterRequest

		if err := c.ShouldBindJSON(&registerRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    http.StatusBadRequest,
				"message": err.Error(),
			})
			return
		}

		userObj, err := userService.GetByEmail(ctx, registerRequest.Email)

		if !errors.Is(err, entity.ErrNotFound) {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":    http.StatusInternalServerError,
				"message": err.Error(),
			})
			return
		}

		if userObj != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    http.StatusBadRequest,
				"message": "User unique",
			})
			return
		}

		userObj, err = userService.Register(ctx, user.CreateUserDTO{
			Name:     registerRequest.Name,
			Email:    registerRequest.Email,
			Password: registerRequest.Password,
		})

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":    http.StatusInternalServerError,
				"message": err.Error(),
			})
			return
		}

		userResp := presenter.UserResponse{
			ID:    userObj.ID,
			Name:  userObj.Name,
			Email: userObj.Email,
		}

		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"message": "success",
			"data": gin.H{
				"user": userResp,
			},
		})
	}
}

func login(userService user.Usecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)

		defer cancel()
		var loginRequest request.LoginRequest

		if err := c.ShouldBindJSON(&loginRequest); err != nil {

			c.JSON(http.StatusBadRequest, gin.H{
				"code":    http.StatusBadRequest,
				"message": err.Error(),
			})
			return
		}

		user, token, err := userService.Login(ctx, loginRequest.Email, loginRequest.Password)

		if errors.Is(err, entity.ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"code":    http.StatusNotFound,
				"message": err.Error(),
			})
			return
		}

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":    http.StatusInternalServerError,
				"message": err.Error(),
			})
		}

		userResp := presenter.UserResponse{
			ID:    user.ID,
			Email: user.Email,
			Name:  user.Name,
		}

		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"message": "success",
			"data": gin.H{

				"user":  userResp,
				"token": token,
			},
		})
	}
}

func NewUserHandler(base *gin.RouterGroup, userService user.Usecase) {
	userGroup := base.Group("users")
	{
		userGroup.POST("login", login(userService))
		userGroup.POST("register", register(userService))
	}
}
