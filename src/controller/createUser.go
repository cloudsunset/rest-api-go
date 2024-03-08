package controller

import (
	"github.com/cloudsunset/api-go/src/configuration/logger"
	"github.com/cloudsunset/api-go/src/configuration/validation"
	"github.com/cloudsunset/api-go/src/controller/model/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CreateUser(c *gin.Context) {

	logger.Info("Create User Controller",
		zap.String("journey", "createUser"),
	)

	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("errors trying to marshal object, error=%s\n", err)

		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	logger.Info("User created successfully",
		zap.String("journey", "createUser"),
	)

	//c.JSON(http.StatusOK, response)
}
