package auth

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"
	"web-framework/configs"
	"web-framework/models/domains"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type IAtuh interface {
	GenerateToken(user *domains.User) (string, error)
	Authenticate() gin.HandlerFunc
}

type Auth struct {
	Secret string
}

func NewAuth(appconfig *configs.AppConfig) IAtuh {
	return &Auth{
		Secret: appconfig.Secret,
	}
}

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func (h *Auth) GenerateToken(user *domains.User) (string, error) {
	// jwt token
	expirationTime := time.Now().Add(time.Minute * 5)

	claims := &Claims{
		Email: user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	fmt.Println(claims)
	fmt.Println("hello from gen")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(h.Secret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (h *Auth) Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		extractToken := func() string {
			bearerToken := c.Request.Header.Get("Authorization")
			strArr := strings.Split(bearerToken, " ")
			if len(strArr) == 2 {
				return strArr[1]
			}
			return ""
		}

		if c.Request.Header["Authorization"] != nil {
			tokenString := extractToken()
			// fmt.Println(tokenString)
			token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
				if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, errors.New("invalid authorized token")
				}
				return []byte(h.Secret), nil
			})

			// fmt.Println("hi ", err)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				c.Abort()
				return
			}
			// fmt.Println("here")
			if token.Valid {
				claims := token.Claims.(jwt.MapClaims)
				c.Set("email", claims["email"])
				c.Next()
			}
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "no authorized header found"})
			c.Abort()
			return
		}
	}
}
