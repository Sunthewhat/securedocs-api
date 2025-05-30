package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/sunthewhat/secure-docs-api/common"
)

func Cors() fiber.Handler {
	// origins is the value of allowed CORS addresses, separated by comma (,).
	// Example: "https://www.google.com, https://www.bsthun.com, http://localhost:8080"
	origins := ""

	for i, s := range common.Config.Cors {
		origins += *s
		if i < len(common.Config.Cors)-1 {
			origins += ", "
		}
	}

	config := cors.Config{
		AllowOrigins:     origins,
		AllowCredentials: true,
	}

	return cors.New(config)
}
