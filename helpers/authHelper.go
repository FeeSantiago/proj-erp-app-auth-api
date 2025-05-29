package helpers

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func CheckUserType(c *gin.Context, role string) (err error) {
	userType := c.GetString("userType")
	err = nil
	if userType == role {
		err = errors.New("unauthorized to access this resource")
		return err
	}
	return err
}

func MatchUserTypeToUserId(c *gin.Context, userId string) (err error) {
	userType := c.GetString("userType")
	uId := c.GetString("userId")
	err = nil
	if userType != "ADMIN" && uId != userId {
		err = errors.New("unauthorized to access this resource")
		return err
	}
	err = CheckUserType(c, userType)
	return err
}
