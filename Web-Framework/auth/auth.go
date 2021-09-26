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
	Given_Name         string `"json": "given_name"`
	Family_Name        string `"json": "family_name"`
	Email              string `json:"email"`
	Company_Id         string `json: "company_id"`
	Product_Codes      string `json: "product_codes"`
	Is_Check_Point_Use bool   `json : "is_check_point_use"`
	jwt.StandardClaims
}

func (h *Auth) GenerateToken(user *domains.User) (string, error) {
	subject := user.Email
	notBefore := time.Now().Add(time.Second * 10)
	expirationTime := notBefore.Add(time.Minute * 10)
	issuer := "localhost"

	claims := &Claims{
		Given_Name:         user.Given_Name,
		Family_Name:        user.Family_Name,
		Email:              user.Email,
		Company_Id:         user.Company_Id,
		Product_Codes:      user.Product_Codes,
		Is_Check_Point_Use: user.Is_Check_Point_Use,
		StandardClaims: jwt.StandardClaims{
			Subject: subject,
			// Id: ,
			NotBefore: notBefore.Unix(),
			ExpiresAt: expirationTime.Unix(),
			Issuer:    issuer,
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
			token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
				if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, errors.New("invalid authorized token")
				}
				return []byte(h.Secret), nil
			})

			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				c.Abort()
				return
			}
			if token.Valid {
				claims := token.Claims.(jwt.MapClaims)
				//fmt.Println(claims)
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
