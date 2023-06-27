package http

import (
	"context"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
)

type service interface {
	Get(ctx context.Context, id string) (json.RawMessage, error)
}

func NewServer(service service) *fiber.App {
	f := fiber.New()

	f.Use(HandleErrors)

	f.Get("/", getHandler(service))

	return f
}
