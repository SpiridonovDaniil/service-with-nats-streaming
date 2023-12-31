package http

import (
	"errors"
	"fmt"
	"net/http"

	"l0/internal/models"

	"github.com/gofiber/fiber/v2"
)

func getHandler(service service) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		id := ctx.Query("key")
		if id == "" {
			ctx.Status(http.StatusBadRequest)
			return fmt.Errorf("[getHandler] search parameters are not specified")
		}

		resp, err := service.Get(ctx.Context(), id)
		if errors.Is(err, models.ErrNotFound) {
			ctx.Status(http.StatusNotFound)
			return fmt.Errorf("[getHandler] %w", err)
		} else if err != nil {
			ctx.Status(http.StatusInternalServerError)
			return fmt.Errorf("[getHandler] %w", err)
		}

		err = ctx.JSON(resp)
		if err != nil {
			ctx.Status(http.StatusInternalServerError)
			return fmt.Errorf("[getHandler] failed to return JSON answer, error: %w", err)
		}
		return nil
	}
}
