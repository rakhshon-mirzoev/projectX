package handlers

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"main.go/pkg/constants"
)

func GenerateToken(userId int64) (string, error) {

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = userId
	claims["exp"] = time.Now().Add(time.Hour * 30).Unix()
	claims["anything"] = "OK"
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(constants.SecretKeyForToken))

}

func GenerateTokenForUnis(uniId int64) (string, error) {

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["uni_id"] = uniId
	claims["exp"] = time.Now().Add(time.Hour * 30).Unix()
	claims["anything"] = "OK"
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(constants.SecretKeyForToken))

}

func GenerateTokenForSchools(schoolId int64) (string, error) {

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["school_id"] = schoolId
	claims["exp"] = time.Now().Add(time.Hour * 30).Unix()
	claims["anything"] = "OK"
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(constants.SecretKeyForToken))

}

func GenerateTokenForAdmins(adminId int64) (string, error) {

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["admin_id"] = adminId
	claims["exp"] = time.Now().Add(time.Hour * 30).Unix()
	claims["anything"] = "ADMIN"
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(constants.SecretKeyForToken))

}

func TokenValid(c *gin.Context) error {
	tokenString := ExtractToken(c)
	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(constants.SecretKeyForToken), nil
	})
	if err != nil {
		fmt.Println("Valid check:", err.Error())
		return err
	}
	return nil
}

func ExtractToken(c *gin.Context) string {
	token := c.Query("token")
	if token != "" {
		fmt.Println("token = ", token)
		return token
	}
	bearerToken := c.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

func ExtractTokenID(c *gin.Context) (uint, error) {
	tokenString := ExtractToken(c)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(constants.SecretKeyForToken), nil
	})
	if err != nil {
		return 0, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		fmt.Println("Token Valid")
		uid, err := strconv.ParseUint(fmt.Sprintf("%.0f", claims["user_id"]), 10, 32)
		if err != nil {
			return 0, err
		}
		fmt.Println("Uid:", uid)
		return uint(uid), nil
	}
	return 0, nil
}
