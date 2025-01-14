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
	ConfigPath string
}

func RunServer(configPath string) error {
	s := server{
		ConfigPath: configPath,
	}
	return s.run()
}

func (s *server) run() error {
	initialConf := LoadConfig(s.ConfigPath)

	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})

	app.Use(cors.New())
	app.Use("/*", func(c *fiber.Ctx) error {
		currentConf := initialConf
		if initialConf.Server.Watch {
			currentConf = LoadConfig(s.ConfigPath)
		}

		var req *requestConfig
		for _, r := range currentConf.Requests {
			if re := regexp.MustCompile(r.Path); re.MatchString(c.Path()) {
				req = &r
				break
			}
		}

		if req == nil {
			slog.Error("request_log",
				slog.String("method", c.Method()),
				slog.String("path", c.Path()),
				slog.String("content_type", c.GetRespHeader(fiber.HeaderContentType)),
			)
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

		slog.Info("request_log",
			slog.String("method", c.Method()),
			slog.String("path", c.Path()),
			slog.String("content_type", c.GetRespHeader(fiber.HeaderContentType)),
		)

		return c.Status(fiber.StatusOK).SendString(res)
	})

	go func() {
		if err := app.Listen(fmt.Sprintf(":%d", initialConf.Server.Port)); err != nil {
			slog.Error(err.Error())
		}
	}()

	slog.Info(fmt.Sprintf("server starts at http://0.0.0.0:%d\n", initialConf.Server.Port))

	ch := make(chan bool)
	<-ch

	return nil
}
