package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"net/http"
	"time"
)

type AuthController struct{}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type JwtOutput struct {
	Token   string    `json:"token"`
	Expires time.Time `json:"expires"`
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (authCon *AuthController) VerifyUser(c *gin.Context) {
	var user User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid username or password",
		})
		return
	}

	if user.Username == "admin" && user.Password == "admin" {

		expirationTime := time.Now().Add(10 * time.Minute)
		cl := &Claims{
			Username: user.Username,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: expirationTime.Unix(),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, cl)
		tokenString, _ := token.SignedString([]byte("SECRET"))

		jwtToken := &JwtOutput{
			Token:   tokenString,
			Expires: expirationTime,
		}

		c.JSON(http.StatusOK, jwtToken)

	}

}
