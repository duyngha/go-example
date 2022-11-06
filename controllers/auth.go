package controllers

import (
	"errors"
	"net/http"
	"time"

	"example.com/m/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// TODO: put the secret key in the .env file
var jwtKey = []byte("supersecretkey")

type JWTClaim struct {
	Name   string `json:"name"`
	KeyID  string `json:"key_id"`
	Secret string `json:"secret"`
	jwt.StandardClaims
}

type TokenRequest struct {
	Name   string `json:"name"`
	KeyID  string `json:"key_id"`
	Secret string `json:"secret"`
}

func GenerateToken(c *gin.Context) {
	var input TokenRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var client models.Client
	if err := models.DB.Where("key_id = ?", input.KeyID).Where("secret = ?", input.Secret).First(&client).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tokenString, err := GenerateJWT(client.Name, client.KeyID, client.Secret)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

func GenerateJWT(name string, key string, secret string) (tokenString string, err error) {
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &JWTClaim{
		Name:   name,
		KeyID:  key,
		Secret: secret,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(jwtKey)
	return
}

func ValidateToken(signedToken string) (err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)

	if err != nil {
		return
	}

	claims, ok := token.Claims.(*JWTClaim)

	if !ok {
		err = errors.New("couldn't parse claims")
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return
	}

	return
}
