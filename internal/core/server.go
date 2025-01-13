package core

import (
	"fmt"
	"log/slog"
	"os"
	"regexp"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type server struct {
	Config *config
}

func RunServer(c *config) error {
	s := server{
		Config: c,
	}
	return s.run()
}

func (s *server) run() error {
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})

	app.Use(cors.New())
	app.Use("/*", func(c *fiber.Ctx) error {
		fmt.Printf("config: %+v\n", s.Config)

		var req *requestConfig
		for _, r := range s.Config.Requests {
			if re := regexp.MustCompile(r.Path); re.MatchString(c.Path()) {
				req = &r
				break
			}
		}

		if req == nil {
			return c.Status(fiber.StatusNotFound).SendString("no response found for the request")
		}

		f, _ := os.ReadFile(req.ResponseFile)
		res := string(f)

		unknownResType := true

		if strings.HasSuffix(req.ResponseFile, "json") {
			c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
			unknownResType = false
		}

		if strings.HasSuffix(req.ResponseFile, ".html") || strings.HasSuffix(req.ResponseFile, ".htm") {
			c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
			unknownResType = false
		}

		if unknownResType {
			c.Set(fiber.HeaderContentType, fiber.MIMETextPlain)
		}

		return c.Status(fiber.StatusOK).SendString(res)
	})

	go func() {
		if err := app.Listen(fmt.Sprintf(":%d", s.Config.Server.Port)); err != nil {
			slog.Error(err.Error())
		}
	}()

	slog.Info(fmt.Sprintf("server starts at http://0.0.0.0:%d\n", s.Config.Server.Port))

	ch := make(chan bool)
	<-ch

	return nil
}
