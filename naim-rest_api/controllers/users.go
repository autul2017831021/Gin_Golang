package controllers

import (
	"net/http"
	"time"
	"web-framework/models"
	"web-framework/models/dtos"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type UserController struct{}

var jwtKey = []byte("secret_key")

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func NewUserController() *UserController {
	return &UserController{}
}

func (h *UserController) Handler(r *gin.Engine) {
	router := r.Group("/users")
	{
		router.GET("/", h.getUser)
		router.POST("/", h.login)
	}

}

func (h *UserController) getUser(ctx *gin.Context) {
	cookie, err := ctx.Request.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			ctx.JSON(http.StatusUnauthorized, err)
			return
		}
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	tokenString := cookie.Value
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims,
		func(t *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			ctx.JSON(http.StatusUnauthorized, err)
			return
		}
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	if !token.Valid {
		ctx.JSON(http.StatusUnauthorized, err)
		return
	}
	userInfo := models.Guser
	ctx.JSON(http.StatusOK, userInfo)
}

func (h *UserController) login(c *gin.Context) {
	loginDto := dtos.LoginDto{}
	err := c.ShouldBindJSON(&loginDto)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	user := models.Guser
	if loginDto.Email != user.Email || loginDto.Password != user.Pass {
		// c.JSON(http.StatusOK, user)
		c.JSON(http.StatusNotFound, gin.H{"error": "Unauthorized"})
		return
	}

	// jwt token
	expirationTime := time.Now().Add(time.Minute * 5)

	claims := &Claims{
		Email: loginDto.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	http.SetCookie(c.Writer,
		&http.Cookie{
			Name:    "token",
			Value:   tokenString,
			Expires: expirationTime,
		})

	c.JSON(http.StatusOK, tokenString)

}
