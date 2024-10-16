package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"main.go/pkg/constants"
)

const (
	authorizationHeader = "Authorization"
	userIDCtx           = "userID"
	userRoleCtx         = "userRole"
)

func CheckUserAuthentication(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)

	if header == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "empty auth header",
		})
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "invalid auth header",
		})
		return
	}

	if len(headerParts[1]) == 0 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "token is empty",
		})
		return
	}

	accessToken := headerParts[1]

	claims, err := ParseTokenAdmin(accessToken)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	if !claims.IsAdmin {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "only the admin can get the list of users",
		})
	}
	c.Next()
}

func GenerateToken(userId int64) (string, error) {

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = userId
	claims["exp"] = time.Now().Add(time.Minute * 300).Unix()
	claims["anything"] = "OK"
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(constants.SecretKeyForToken))

}

func GenerateTokenForUnis(uniId int64) (string, error) {

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["uni_id"] = uniId
	claims["exp"] = time.Now().Add(time.Minute * 300).Unix()
	claims["anything"] = "OK"
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(constants.SecretKeyForToken))

}

func GenerateTokenForSchools(schoolId int64) (string, error) {

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["school_id"] = schoolId
	claims["exp"] = time.Now().Add(time.Minute * 300).Unix()
	claims["anything"] = "OK"
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(constants.SecretKeyForToken))

}

func GenerateTokenForAdmins(adminId int64) (string, error) {

	claims := CustomClaims{
		Authorized: true,
		AdminId:    adminId,
		IsAdmin:    true,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 60).Unix(),
		},
	}
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

type CustomClaims struct {
	Authorized bool
	AdminId    int64
	IsAdmin    bool
	jwt.StandardClaims
}

func ParseTokenAdmin(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Проверяем метод подписи токена
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(constants.SecretKeyForToken), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}
