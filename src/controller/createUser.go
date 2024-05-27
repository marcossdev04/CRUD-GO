package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/marcossdev04/crud-go/src/configuration/logger"
	"github.com/marcossdev04/crud-go/src/controller/model/request"
	"github.com/marcossdev04/crud-go/src/model"
	"github.com/marcossdev04/crud-go/src/model/service"
	"github.com/marcossdev04/crud-go/src/validation"
	"go.uber.org/zap"
)

func CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser controler",
		zap.String("journey", "createUser"),
	)
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err)
		errRest := validation.ValidateUserError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	domain := model.NewUserDomain(
		userRequest.Email,
		userRequest.Password,
		userRequest.Name,
		userRequest.Age,
	)

	service := service.NewUserDomainService()

	if err := service.CreateUser(domain); err != nil {
		c.JSON(err.Code, err)
	}

	logger.Info("User Created successfully",
		zap.String("journey", "createUser"))
	c.String(http.StatusOK, "")
}
