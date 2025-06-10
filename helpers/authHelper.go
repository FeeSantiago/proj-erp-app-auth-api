package helpers

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
)

func CheckUserType(c *gin.Context, role string) (err error) {
	userType := c.GetString("userType")
	err = nil
	if userType != role {
		err = errors.New("unauthorized to access this resource")
		return err
	}
	return err
}

func MatchUserTypeToUserId(c *gin.Context, userId string) (err error) {
	userType := c.GetString("userType")
	uId := c.GetString("userId")
	fmt.Println("INFO userType:", userType)
	fmt.Println("INFO userId:", uId)
	err = nil
	if userType != "ADMIN" && uId != userId {
		err = errors.New("unauthorized to access this resource")
		fmt.Println(err)
		return err
	}
	err = CheckUserType(c, userType)
	return err
}
