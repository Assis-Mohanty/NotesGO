package middleware

import (
	"net/http"
	"strings"

	"github.com/Assis-Mohanty/notes/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func JWTProtected() fiber.Handler{
	return func (c *fiber.Ctx) error {
		authHeader:=c.Get("Authorization")
		if authHeader==""{
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"Error":"Invalid token"})
		}
		
		parts := strings.Split(authHeader," ")
		if len(parts)!=2 || parts[0]!="Bearer"{
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"Error":"invalid token format"})
		}
		token,err:=utils.VerifyJWT(parts[1],false)
		if err != nil || !token.Valid {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"error": "invalid or expired token",
			})
		}
		claims := token.Claims.(jwt.MapClaims)
		c.Locals("user_id",uint(claims["user_id"].(float64)))
		return c.Next()
	}
}