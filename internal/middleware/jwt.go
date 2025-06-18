package middleware

import (
	"boilerplate/pkg/exception"
	"boilerplate/pkg/utils"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func EmployeeAuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		init := exception.InitException(c, initData.Conf, initData.Log)

		authorizationHeader := c.Get("Authorization")
		if !strings.Contains(authorizationHeader, "Bearer") {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"status":  "error",
				"message": "Invalid token format",
			})
		}
		accessToken := strings.Replace(authorizationHeader, "Bearer ", "", -1)
		claims, err := utils.CheckAccessToken(init.Conf, accessToken)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"status":  "error",
				"message": "Token expired",
			})
		}

		employeeName := claims["employee_name"].(string)
		c.Locals("employee_name", employeeName)

		return c.Next()
	}
}

func AdminAuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		init := exception.InitException(c, initData.Conf, initData.Log)

		authorizationHeader := c.Get("Authorization")
		if !strings.Contains(authorizationHeader, "Bearer") {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"status":  "error",
				"message": "Invalid token format",
			})
		}
		accessToken := strings.Replace(authorizationHeader, "Bearer ", "", -1)
		claims, err := utils.CheckAccessToken(init.Conf, accessToken)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"status":  "error",
				"message": "Token expired" + err.Error(),
			})
		}

		employeeName := claims["employee_name"].(string)
		roleName := claims["role_name"].(string)
		if roleName != "admin" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"status":  "error",
				"message": "Role unauthorized",
			})
		}
		c.Locals("employee_name", employeeName)
		c.Locals("role_name", roleName)

		return c.Next()
	}
}
