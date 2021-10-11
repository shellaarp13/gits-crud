package middleware

import (
	"log"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type jwtCustomClaims struct {
	UserId int    `json:"user_id"`
	Name   string `json:"name"`
	Role   string `json:"role"`
	jwt.StandardClaims
}

type ConfigJWT struct {
	SecretJWT string
	ExpiredIn int
}

func (jwtConf *ConfigJWT) Init() middleware.JWTConfig {
	return middleware.JWTConfig{
		Claims:     &jwtCustomClaims{},
		SigningKey: []byte(jwtConf.SecretJWT),
	}
}

//generate new token
func (jwtConf *ConfigJWT) GenerateToken(id int, name string, role string, status string) string {
	claims := &jwtCustomClaims{
		id,
		name,
		role,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * time.Duration(int64(jwtConf.ExpiredIn))).Unix(),
		},
	}
	log.Println("JWT CONFIG LOOKa AT DIS", jwtConf)
	//membuat token dari claims yang isinya data" tersebut
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(jwtConf.SecretJWT))
	log.Println("INI TOKEN", t)
	exceptions.PanicIfError(err)

	return t
}

func GetUserId(c echo.Context) int {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*jwtCustomClaims)
	id := claims.UserId
	return id
}

func GetUserName(c echo.Context) string {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*jwtCustomClaims)
	name := claims.Name
	return name
}

func GetUserRole(c echo.Context) string {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*jwtCustomClaims)
	role := claims.Role
	return role
}
